package cliente

import (
	"fmt"
	"go/mydelivery/shared/errs"
	"go/mydelivery/shared/validator"
)

type cliente struct {
	id       int64
	nome     string
	email    string
	telefone string
}

func New(nome, email, telefone string) Cliente {
	return &cliente{0, nome, email, telefone}
}

func NewWithID(id int64, nome, email, telefone string) Cliente {
	return &cliente{id, nome, email, telefone}
}

func (c *cliente) ID() int64 { return c.id }

func (c *cliente) Nome() string { return c.nome }

func (c *cliente) Email() string { return c.email }

func (c *cliente) Telefone() string { return c.telefone }

func (c *cliente) String() string {
	return fmt.Sprintf("\nID: %d\nNome: %s\nEmail: %s\nTelefone: %s\n",
		c.id, c.nome, c.email, c.telefone)
}

func (c *cliente) Validar() error {
	v := validator.New()
	v.CheckBlank("nome", c.nome)
	v.CheckEmail("email", c.email)
	v.CheckBlank("telefone", c.telefone)

	if v.HasErrors() {
		return errs.NewValidationError(v.Errors)
	}

	return nil
}
