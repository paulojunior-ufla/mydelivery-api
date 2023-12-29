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

func (r *entregaRepositoryDB) Todos() ([]Entrega, error) {
	query := `
		SELECT e.id, e.taxa, e.status, e.data_pedido, e.dest_nome, e.dest_endereco, c.id, c.nome, c.email, c.telefone
		FROM entrega AS e INNER JOIN cliente AS c ON e.cliente_id = c.id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errs.NewUnexpectedError(err)
	}
	defer rows.Close()

	entregas := []Entrega{}
	for rows.Next() {
		var c cliente
		var d destinatario
		var e entrega

		err := rows.Scan(&e.id, &e.taxa, &e.status, &e.dataPedido, &d.nome, &d.endereco, &c.id, &c.nome, &c.email, &c.telefone)
		if err != nil {
			return nil, errs.NewUnexpectedError(err)
		}

		e.cliente = &c
		e.destinatario = &d

		entregas = append(entregas, &e)
	}

	if err = rows.Err(); err != nil {
		return nil, errs.NewUnexpectedError(err)
	}

	return entregas, nil
}

func (r *entregaRepositoryDB) Salvar(e Entrega) (int64, error) {
	query := `
		INSERT INTO entrega(cliente_id, taxa, status, data_pedido, dest_nome, dest_endereco) 
		VALUES(?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query,
		e.Cliente().ID(), e.Taxa(), e.Status(), e.DataPedido(), e.Destinatario().Nome(), e.Destinatario().Endereco())

	if err != nil {
		return 0, errs.NewUnexpectedError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errs.NewUnexpectedError(err)
	}

	return id, nil
}

func (r *entregaRepositoryDB) Atualizar(e Entrega) error {
	query := "UPDATE entrega SET cliente_id = ?, taxa = ?, status = ?, data_pedido = ?, dest_nome = ?, dest_endereco = ? WHERE id = ?"

	result, err := r.db.Exec(query, e.Cliente().ID(), e.Taxa(), e.Status(), e.DataPedido(), e.Destinatario().Nome(), e.Destinatario().Endereco(), e.ID())
	if err != nil {
		return errs.NewUnexpectedError(err)
	}

	linhasAfetadas, err := result.RowsAffected()
	if err != nil {
		return errs.NewUnexpectedError(err)
	}

	if linhasAfetadas == 0 {
		return errs.NewNotFoundError("entrega não encontrada")
	}

	return nil
}

func (r *entregaRepositoryDB) ObterPorID(id int64) (Entrega, error) {
	query := `
		SELECT e.id, e.taxa, e.status, e.data_pedido, e.dest_nome, e.dest_endereco, c.id, c.nome, c.email, c.telefone
		FROM entrega AS e INNER JOIN cliente AS c ON e.cliente_id = c.id
		WHERE e.id = ?`

	var c cliente
	var d destinatario
	var e entrega

	err := r.db.QueryRow(query, id).
		Scan(&e.id, &e.taxa, &e.status, &e.dataPedido, &d.nome, &d.endereco, &c.id, &c.nome, &c.email, &c.telefone)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.NewUnexpectedError(err)
	}

	e.cliente = &c
	e.destinatario = &d

	return &e, nil
}
