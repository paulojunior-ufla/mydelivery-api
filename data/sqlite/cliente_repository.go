package sqlite

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

	entityList := []*clienteEntity{}
	for rows.Next() {
		var entity clienteEntity
		err := rows.Scan(&entity.id, &entity.nome, &entity.email, &entity.telefone)
		if err != nil {
			return nil, errs.NewUnexpectedError(err)
		}
		entityList = append(entityList, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, errs.NewUnexpectedError(err)
	}

	return toClienteList(entityList), nil
}

func (r *clienteRepository) ObterPorID(id int64) (cliente.Cliente, error) {
	query := "SELECT id, nome, email, telefone FROM cliente WHERE id = ?"

	var entity clienteEntity
	err := r.db.QueryRow(query, id).Scan(&entity.id, &entity.nome, &entity.email, &entity.telefone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("cliente não encontrado")
		}
		return nil, errs.NewUnexpectedError(err)
	}

	return toCliente(&entity), nil
}

func (r *clienteRepository) ObterPorEmail(email string) (cliente.Cliente, error) {
	query := "SELECT id, nome, email, telefone FROM cliente WHERE email = ?"

	var entity clienteEntity
	err := r.db.QueryRow(query, email).Scan(&entity.id, &entity.nome, &entity.email, &entity.telefone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.NewUnexpectedError(err)
	}

	return toCliente(&entity), nil
}

func (r *clienteRepository) Salvar(c cliente.Cliente) (cliente.Cliente, error) {
	query := "INSERT INTO cliente(nome, email, telefone) VALUES(?, ?, ?)"

	result, err := r.db.Exec(query, c.Nome(), c.Email(), c.Telefone())
	if err != nil {
		return nil, errs.NewUnexpectedError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errs.NewUnexpectedError(err)
	}

	return cliente.NewWithID(id, c.Nome(), c.Email(), c.Telefone()), nil
}

func (r *clienteRepository) Atualizar(c cliente.Cliente) (cliente.Cliente, error) {
	query := "UPDATE cliente SET nome = ?, email = ?, telefone = ? WHERE id = ?"

	result, err := r.db.Exec(query, c.Nome(), c.Email(), c.Telefone(), c.ID())
	if err != nil {
		return nil, errs.NewUnexpectedError(err)
	}

	linhasAfetadas, err := result.RowsAffected()
	if err != nil {
		return nil, errs.NewUnexpectedError(err)
	}

	if linhasAfetadas == 0 {
		return nil, errs.NewNotFoundError("cliente não encontrado")
	}

	return cliente.NewWithID(c.ID(), c.Nome(), c.Email(), c.Telefone()), nil
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

func toClienteList(clienteEntityList []*clienteEntity) []cliente.Cliente {
	clientes := []cliente.Cliente{}
	for _, entity := range clienteEntityList {
		clientes = append(clientes, toCliente(entity))
	}
	return clientes
}

func toCliente(entity *clienteEntity) cliente.Cliente {
	return cliente.NewWithID(entity.id, entity.nome, entity.email, entity.telefone)
}
