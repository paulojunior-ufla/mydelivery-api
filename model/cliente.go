package model

import (
	"go/mydelivery/shared/errs"
	"go/mydelivery/shared/validator"
)

type ClienteRepository interface {
	Todos() ([]Cliente, error)
	ObterPorID(int64) (Cliente, error)
	ObterPorEmail(string) (Cliente, error)
	Salvar(Cliente) (int64, error)
	Atualizar(Cliente) error
	Excluir(int64) error
}

type Cliente interface {
	ID() int64
	Nome() string
	Email() string
	Telefone() string
}

type cliente struct {
	id       int64
	nome     string
	email    string
	telefone string
}

func (c *cliente) ID() int64        { return c.id }
func (c *cliente) Nome() string     { return c.nome }
func (c *cliente) Email() string    { return c.email }
func (c *cliente) Telefone() string { return c.telefone }

type ClienteBuilder struct {
	cliente *cliente
}

func NewCliente() *ClienteBuilder {
	return &ClienteBuilder{
		cliente: &cliente{},
	}
}

func (b *ClienteBuilder) SetID(id int64) *ClienteBuilder {
	b.cliente.id = id
	return b
}

func (b *ClienteBuilder) SetNome(nome string) *ClienteBuilder {
	b.cliente.nome = nome
	return b
}

func (b *ClienteBuilder) SetEmail(email string) *ClienteBuilder {
	b.cliente.email = email
	return b
}

func (b *ClienteBuilder) SetTelefone(telefone string) *ClienteBuilder {
	b.cliente.telefone = telefone
	return b
}

func (b *ClienteBuilder) Build() (Cliente, error) {
	v := validator.New()

	v.CheckBlank("nome do cliente", b.cliente.nome)
	v.CheckEmail("email do cliente", b.cliente.email)
	v.CheckBlank("telefone do cliente", b.cliente.telefone)

	if v.HasErrors() {
		return nil, errs.NewValidationError(v.Errors)
	}

	return b.cliente, nil
}
