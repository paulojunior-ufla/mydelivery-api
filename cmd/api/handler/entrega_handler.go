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
	finalizaEntregaSrv entrega.FinalizaEntregaService
}

func NewEntregaHandler(
	entregaRepo model.EntregaRepository,
	solicitaEntregaSrv entrega.SolicitaEntregaService,
	finalizaEntregaSrv entrega.FinalizaEntregaService) *entregaHandler {
	return &entregaHandler{entregaRepo, solicitaEntregaSrv, finalizaEntregaSrv}
}

func (h *entregaHandler) InitRoutes(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/entregas", h.Listar)
	router.HandlerFunc(http.MethodPost, "/entregas", h.Solicitar)
	router.HandlerFunc(http.MethodGet, "/entregas/:id", h.BuscarPorID)
	router.HandlerFunc(http.MethodPut, "/entregas/:id/finalizar", h.Finalizar)
}

func (h *entregaHandler) Listar(w http.ResponseWriter, r *http.Request) {
	entregas, err := h.entregaRepo.Todos()
	if err != nil {
		handleError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, entrega.ToEntregaResponseCollection(entregas))
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

	writeJSON(w, http.StatusOK, entrega.ToEntregaResponse(e))
}

func (h *entregaHandler) Finalizar(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = h.finalizaEntregaSrv.Finalizar(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
