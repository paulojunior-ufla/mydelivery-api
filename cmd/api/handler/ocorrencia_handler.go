package handler

import (
	"go/mydelivery/model"
	"go/mydelivery/service/ocorrencia"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ocorrenciaHandler struct {
	ocorrenciasRepo       model.OcorrenciaRepository
	registrarOcorrenciSrv ocorrencia.RegistraOcorrenciaService
}

func NewOcorrenciaHandler(
	ocorrenciasRepo model.OcorrenciaRepository,
	registrarOcorrenciSrv ocorrencia.RegistraOcorrenciaService) *ocorrenciaHandler {
	return &ocorrenciaHandler{ocorrenciasRepo, registrarOcorrenciSrv}
}

func (h *ocorrenciaHandler) InitRoutes(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/entregas/:id/ocorrencias", h.Listar)
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

func (h *ocorrenciaHandler) Listar(w http.ResponseWriter, r *http.Request) {
	ocorrencias, err := h.ocorrenciasRepo.Todos()
	if err != nil {
		handleError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, ocorrencia.ToRegistraOcorrenciaResponseCollection(ocorrencias))
}
