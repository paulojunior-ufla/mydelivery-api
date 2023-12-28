package handler

import (
	"fmt"
	"go/mydelivery/shared/errs"
	"log/slog"
	"net/http"
)

type errorMessage struct {
	Message string   `json:"error"`
	Details []string `json:"details,omitempty"`
}

func logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	slog.Error(err.Error(), "method", method, "uri", uri)
}

func serverError(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)
	writeJSON(w, http.StatusInternalServerError, errorMessage{
		Message: "erro interno do sistema",
	})
}

func appError(w http.ResponseWriter, r *http.Request, err *errs.Err) {
	switch err.ErrType {
	case errs.VALIDATION_ERROR:
		writeJSON(w, http.StatusUnprocessableEntity, errorMessage{
			Message: err.Message,
			Details: err.Details,
		})
	case errs.UNEXPECTED_ERROR:
		serverError(w, r, err.Cause)
	case errs.NOT_FOUND_ERROR:
		writeJSON(w, http.StatusNotFound, errorMessage{
			Message: err.Message,
		})
	case errs.BAD_REQUEST_ERROR:
		logError(r, err.Cause)
		writeJSON(w, http.StatusBadRequest, errorMessage{
			Message: err.Message,
		})
	case errs.CONFLICT_ERROR:
		writeJSON(w, http.StatusUnprocessableEntity, errorMessage{
			Message: err.Message,
		})
	default:
		serverError(w, r, fmt.Errorf("tipo de erro não suportado: %s", err.ErrType))
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	switch e := err.(type) {
	case *errs.Err:
		appError(w, r, e)
	default:
		serverError(w, r, err)
	}
}
