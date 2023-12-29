package model

import (
	"database/sql"
	"go/mydelivery/shared/errs"
)

type ocorrenciaRepositoryDB struct {
	db *sql.DB
}

func NewOcorrenciaRepositoryDB(db *sql.DB) *ocorrenciaRepositoryDB {
	return &ocorrenciaRepositoryDB{db}
}

func (r *ocorrenciaRepositoryDB) Salvar(o Ocorrencia) (int64, error) {
	query := `
		INSERT INTO ocorrencia(entrega_id, descricao, data_registro) 
		VALUES(?, ?, ?)`

	result, err := r.db.Exec(query,
		o.Entrega().ID(), o.Descricao(), o.DataRegistro())

	if err != nil {
		return 0, errs.NewUnexpectedError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errs.NewUnexpectedError(err)
	}

	return id, nil
}
