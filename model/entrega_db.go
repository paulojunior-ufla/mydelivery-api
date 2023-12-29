package model

import (
	"database/sql"
	"errors"
	"go/mydelivery/shared/errs"
)

type entregaRepositoryDB struct {
	db *sql.DB
}

func NewEntregaRepositoryDB(db *sql.DB) *entregaRepositoryDB {
	return &entregaRepositoryDB{db}
}
func (r *entregaRepositoryDB) Salvar(e Entrega) (int64, error) {
	query := "INSERT INTO entrega(cliente_id, taxa, status, data_pedido) VALUES(?, ?, ?, ?)"

	result, err := r.db.Exec(query, e.Cliente().ID(), e.Taxa(), e.Status(), e.DataPedido())
	if err != nil {
		return 0, errs.NewUnexpectedError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errs.NewUnexpectedError(err)
	}

	return id, nil
}

func (r *entregaRepositoryDB) ObterPorID(id int64) (Entrega, error) {
	query := `
		SELECT e.id, e.taxa, e.status, e.data_pedido, c.id, c.nome, c.email, c.telefone  
		FROM entrega AS e INNER JOIN cliente AS c ON e.cliente_id = c.id
		WHERE e.id = ?`

	var c cliente
	var e entrega

	err := r.db.QueryRow(query, id).
		Scan(&e.id, &e.taxa, &e.status, &e.dataPedido, &c.id, &c.nome, &c.email, &c.telefone)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.NewUnexpectedError(err)
	}

	e.cliente = &c

	return &e, nil
}
