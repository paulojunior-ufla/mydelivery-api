package data

import (
	"database/sql"
	"errors"
	"go/mydelivery/domain/cliente"
	"go/mydelivery/shared/errs"
)

type clienteEntity struct {
	id       int64
	nome     string
	email    string
	telefone string
}

type clienteRepository struct {
	db *sql.DB
}

func NewClienteRepository(db *sql.DB) *clienteRepository {
	return &clienteRepository{db}
}

func (r *clienteRepository) Todos() ([]cliente.Cliente, error) {
	query := "SELECT id, nome, email, telefone FROM cliente"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errs.NewUnexpectedError(err)
	}
	defer rows.Close()

	entities := []clienteEntity{}
	for rows.Next() {
		var e clienteEntity
		err := rows.Scan(&e.id, &e.nome, &e.email, &e.telefone)
		if err != nil {
			return nil, errs.NewUnexpectedError(err)
		}
		entities = append(entities, e)
	}

	if err = rows.Err(); err != nil {
		return nil, errs.NewUnexpectedError(err)
	}

	return toClienteCollection(entities), nil
}

func (r *clienteRepository) ObterPorID(id int64) (cliente.Cliente, error) {
	query := "SELECT id, nome, email, telefone FROM cliente WHERE id = ?"

	var e clienteEntity
	err := r.db.QueryRow(query, id).Scan(&e.id, &e.nome, &e.email, &e.telefone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.NewUnexpectedError(err)
	}

	return toCliente(e), nil
}

func (r *clienteRepository) ObterPorEmail(email string) (cliente.Cliente, error) {
	query := "SELECT id, nome, email, telefone FROM cliente WHERE email = ?"

	var e clienteEntity
	err := r.db.QueryRow(query, email).Scan(&e.id, &e.nome, &e.email, &e.telefone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.NewUnexpectedError(err)
	}

	return toCliente(e), nil
}

func (r *clienteRepository) Salvar(c cliente.Cliente) (int64, error) {
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

func (r *clienteRepository) Atualizar(c cliente.Cliente) error {
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

func (r *clienteRepository) Excluir(id int64) error {
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

func toClienteCollection(entities []clienteEntity) []cliente.Cliente {
	clientes := []cliente.Cliente{}
	for _, e := range entities {
		clientes = append(clientes, toCliente(e))
	}
	return clientes
}

func toCliente(e clienteEntity) cliente.Cliente {
	cliente, _ := cliente.New().
		SetID(e.id).
		SetNome(e.nome).
		SetEmail(e.email).
		SetTelefone(e.telefone).
		Build()

	return cliente
}
