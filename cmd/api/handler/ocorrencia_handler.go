package handler

import (
	"go/mydelivery/service/ocorrencia"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ocorrenciaHandler struct {
	registrarOcorrenciSrv ocorrencia.RegistraOcorrenciaService
}

func NewOcorrenciaHandler(
	registrarOcorrenciSrv ocorrencia.RegistraOcorrenciaService) *ocorrenciaHandler {
	return &ocorrenciaHandler{registrarOcorrenciSrv}
}

func (h *ocorrenciaHandler) InitRoutes(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/entregas/:id/ocorrencias", h.RegistrarOcorrencia)
}

func (h *ocorrenciaHandler) RegistrarOcorrencia(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		handleError(w, r, err)
		return
	}

	var input ocorrencia.RegistraOcorrenciaRequest
	err = readJSON(r, &input)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err := h.registrarOcorrenciSrv.Registrar(id, input)
	if err != nil {
		handleError(w, r, err)
		return
	}

	writeJSON(w, http.StatusCreated, o)
}
