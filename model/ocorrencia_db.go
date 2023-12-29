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

func (r *ocorrenciaRepositoryDB) Todos() ([]Ocorrencia, error) {
	query := `
		SELECT o.id, o.descricao, o.data_registro, e.id, e.taxa, e.status, e.data_pedido, e.dest_nome, e.dest_endereco, c.id, c.nome, c.email, c.telefone
		FROM ocorrencia AS o 
		INNER JOIN entrega AS e ON o.entrega_id = e.id
		INNER JOIN cliente AS c ON e.cliente_id = c.id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errs.NewUnexpectedError(err)
	}
	defer rows.Close()

	ocorrencias := []Ocorrencia{}
	for rows.Next() {
		var o ocorrencia
		var c cliente
		var d destinatario
		var e entrega

		err := rows.Scan(&o.id, &o.descricao, &o.dataRegistro, &e.id, &e.taxa, &e.status, &e.dataPedido, &d.nome, &d.endereco, &c.id, &c.nome, &c.email, &c.telefone)
		if err != nil {
			return nil, errs.NewUnexpectedError(err)
		}

		e.cliente = &c
		e.destinatario = &d
		o.entrega = &e

		ocorrencias = append(ocorrencias, &o)
	}

	if err = rows.Err(); err != nil {
		return nil, errs.NewUnexpectedError(err)
	}

	return ocorrencias, nil
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
