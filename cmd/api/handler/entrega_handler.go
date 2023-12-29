package handler

import (
	"go/mydelivery/service/entrega"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type entregaHandler struct {
	solicitaEntregaSrv entrega.SolicitaEntregaService
}

func NewEntregaHandler(solicitaEntregaSrv entrega.SolicitaEntregaService) *entregaHandler {
	return &entregaHandler{solicitaEntregaSrv}
}

func (h *entregaHandler) InitRoutes(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/entregas", h.Solicitar)
}

func (h *entregaHandler) Solicitar(w http.ResponseWriter, r *http.Request) {
	var input entrega.SolicitaEntregaRequest
	err := readJSON(r, &input)
	if err != nil {
		handleError(w, r, err)
		return
	}

	c, err := h.solicitaEntregaSrv.Solicitar(input)
	if err != nil {
		handleError(w, r, err)
		return
	}

	writeJSON(w, http.StatusCreated, c)
}
