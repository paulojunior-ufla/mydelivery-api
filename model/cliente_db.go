package model

import (
	"database/sql"
	"errors"
	"go/mydelivery/shared/errs"
)

type clienteRepositoryDB struct {
	db *sql.DB
}

func NewClienteRepositoryDB(db *sql.DB) *clienteRepositoryDB {
	return &clienteRepositoryDB{db}
}

func (r *clienteRepositoryDB) Todos() ([]Cliente, error) {
	query := "SELECT id, nome, email, telefone FROM cliente"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errs.NewUnexpectedError(err)
	}
	defer rows.Close()

	clientes := []Cliente{}
	for rows.Next() {
		var c cliente
		err := rows.Scan(&c.id, &c.nome, &c.email, &c.telefone)
		if err != nil {
			return nil, errs.NewUnexpectedError(err)
		}
		clientes = append(clientes, &c)
	}

	if err = rows.Err(); err != nil {
		return nil, errs.NewUnexpectedError(err)
	}

	return clientes, nil
}

func (r *clienteRepositoryDB) ObterPorID(id int64) (Cliente, error) {
	query := "SELECT id, nome, email, telefone FROM cliente WHERE id = ?"

	var c cliente
	err := r.db.QueryRow(query, id).Scan(&c.id, &c.nome, &c.email, &c.telefone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("cliente não encontrado")
		}
		return nil, errs.NewUnexpectedError(err)
	}

	return &c, nil
}

func (r *clienteRepositoryDB) ObterPorEmail(email string) (Cliente, error) {
	query := "SELECT id, nome, email, telefone FROM cliente WHERE email = ?"

	var c cliente
	err := r.db.QueryRow(query, email).Scan(&c.id, &c.nome, &c.email, &c.telefone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.NewUnexpectedError(err)
	}

	return &c, nil
}

func (r *clienteRepositoryDB) Salvar(c Cliente) (int64, error) {
	query := "INSERT INTO cliente(nome, email, telefone) VALUES(?, ?, ?)"

	result, err := r.db.Exec(query, c.Nome(), c.Email(), c.Telefone())
	if err != nil {
		return 0, errs.NewUnexpectedError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errs.NewUnexpectedError(err)
	}

	return id, nil
}

func (r *clienteRepositoryDB) Atualizar(c Cliente) error {
	query := "UPDATE cliente SET nome = ?, email = ?, telefone = ? WHERE id = ?"

	result, err := r.db.Exec(query, c.Nome(), c.Email(), c.Telefone(), c.ID())
	if err != nil {
		return errs.NewUnexpectedError(err)
	}

	linhasAfetadas, err := result.RowsAffected()
	if err != nil {
		return errs.NewUnexpectedError(err)
	}

	if linhasAfetadas == 0 {
		return errs.NewNotFoundError("cliente não encontrado")
	}

	return nil
}

func (r *clienteRepositoryDB) Excluir(id int64) error {
	query := "DELETE FROM cliente WHERE id = ?"

	result, err := r.db.Exec(query, id)
	if err != nil {
		return errs.NewUnexpectedError(err)
	}

	linhasAfetadas, err := result.RowsAffected()
	if err != nil {
		return errs.NewUnexpectedError(err)
	}

	if linhasAfetadas == 0 {
		return errs.NewNotFoundError("cliente não encontrado")
	}

	return nil
}
