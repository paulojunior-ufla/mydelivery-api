package cliente

import (
	"go/mydelivery/shared/errs"
	"go/mydelivery/shared/validator"
)

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

func New() *ClienteBuilder {
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

	v.CheckBlank("nome", b.cliente.nome)
	v.CheckEmail("email", b.cliente.email)
	v.CheckBlank("telefone", b.cliente.telefone)

	if v.HasErrors() {
		return nil, errs.NewValidationError(v.Errors)
	}

	return b.cliente, nil
}
