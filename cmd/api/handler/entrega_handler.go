package handler

import (
	"go/mydelivery/model"
	"go/mydelivery/service/entrega"
	"go/mydelivery/shared/errs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type entregaHandler struct {
	entregaRepo        model.EntregaRepository
	solicitaEntregaSrv entrega.SolicitaEntregaService
}

func NewEntregaHandler(
	entregaRepo model.EntregaRepository,
	solicitaEntregaSrv entrega.SolicitaEntregaService) *entregaHandler {
	return &entregaHandler{entregaRepo, solicitaEntregaSrv}
}

func (h *entregaHandler) InitRoutes(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/entregas", h.Solicitar)
	router.HandlerFunc(http.MethodGet, "/entregas/:id", h.BuscarPorID)
}

func (h *entregaHandler) Solicitar(w http.ResponseWriter, r *http.Request) {
	var input entrega.SolicitaEntregaRequest
	err := readJSON(r, &input)
	if err != nil {
		handleError(w, r, err)
		return
	}

	e, err := h.solicitaEntregaSrv.Solicitar(input)
	if err != nil {
		handleError(w, r, err)
		return
	}

	writeJSON(w, http.StatusCreated, e)
}

func (h *entregaHandler) BuscarPorID(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		handleError(w, r, err)
		return
	}

	e, err := h.entregaRepo.ObterPorID(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	if e == nil {
		handleError(w, r, errs.NewNotFoundError("entrega não encontrada"))
		return
	}

	writeJSON(w, http.StatusOK, entrega.ToSolicitaEntregaResponse(e))
}
