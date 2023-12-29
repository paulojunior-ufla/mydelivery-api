package handler

import (
	"go/mydelivery/model"
	"go/mydelivery/service/cliente"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type clienteHandler struct {
	clienteRepo model.ClienteRepository
	clienteSrv  cliente.CatalogoService
}

func NewClienteHandler(clienteRepo model.ClienteRepository, clienteSrv cliente.CatalogoService) *clienteHandler {
	return &clienteHandler{clienteRepo, clienteSrv}
}

func (h *clienteHandler) InitRoutes(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/clientes", h.Salvar)
	router.HandlerFunc(http.MethodPut, "/clientes/:id", h.Atualizar)
	router.HandlerFunc(http.MethodGet, "/clientes", h.Listar)
	router.HandlerFunc(http.MethodGet, "/clientes/:id", h.BuscarPorID)
	router.HandlerFunc(http.MethodDelete, "/clientes/:id", h.Excluir)
}

func (h *clienteHandler) Listar(w http.ResponseWriter, r *http.Request) {
	clientes, err := h.clienteRepo.Todos()
	if err != nil {
		handleError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, cliente.ToClienteResponseCollection(clientes))
}

func (h *clienteHandler) BuscarPorID(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		handleError(w, r, err)
		return
	}

	c, err := h.clienteRepo.ObterPorID(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, cliente.ToClienteResponse(c))
}

func (h *clienteHandler) Salvar(w http.ResponseWriter, r *http.Request) {
	var input cliente.CatalogoClienteRequest
	err := readJSON(r, &input)
	if err != nil {
		handleError(w, r, err)
		return
	}

	c, err := h.clienteSrv.Salvar(input)
	if err != nil {
		handleError(w, r, err)
		return
	}

	writeJSON(w, http.StatusCreated, c)
}

func (h *clienteHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		handleError(w, r, err)
		return
	}

	var input cliente.CatalogoClienteRequest
	err = readJSON(r, &input)
	if err != nil {
		handleError(w, r, err)
		return
	}

	c, err := h.clienteSrv.Atualizar(id, input)
	if err != nil {
		handleError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, c)
}

func (h *clienteHandler) Excluir(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = h.clienteRepo.Excluir(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
