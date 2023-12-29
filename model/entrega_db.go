package model

import (
	"database/sql"
	"go/mydelivery/shared/errs"
)

type entregaRepositoryDB struct {
	db *sql.DB
}

func NewEntregaRepositoryDB(db *sql.DB) *entregaRepositoryDB {
	return &entregaRepositoryDB{db}
}
func (r *entregaRepositoryDB) Salvar(e Entrega) (int64, error) {
	query := "INSERT INTO entrega(cliente_id, taxa, status) VALUES(?, ?, ?)"

	result, err := r.db.Exec(query, e.Cliente().ID(), e.Taxa(), e.Status())
	if err != nil {
		return 0, errs.NewUnexpectedError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errs.NewUnexpectedError(err)
	}

	return id, nil
}
