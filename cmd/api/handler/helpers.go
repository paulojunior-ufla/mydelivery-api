package handler

import (
	"encoding/json"
	"go/mydelivery/shared/errs"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func readJSON(r *http.Request, data any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&data)
	if err != nil {
		return errs.NewBadRequestError(err)
	}

	return nil
}

func readIDParam(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(httprouter.ParamsFromContext(r.Context()).ByName("id"), 10, 64)
	if err != nil {
		return 0, errs.NewBadRequestError(err)
	}
	return id, nil
}
