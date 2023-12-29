package model

import (
	"go/mydelivery/shared/errs"
	"go/mydelivery/shared/validator"
)

type Destinatario interface {
	Nome() string
	Endereco() string
}

type destinatario struct {
	nome     string
	endereco string
}

func (d *destinatario) Nome() string     { return d.nome }
func (d *destinatario) Endereco() string { return d.endereco }

type destinatarioBuilder struct {
	destinatario *destinatario
}

func NewDestinatario() *destinatarioBuilder {
	return &destinatarioBuilder{
		destinatario: &destinatario{},
	}
}

func (b *destinatarioBuilder) SetNome(nome string) *destinatarioBuilder {
	b.destinatario.nome = nome
	return b
}

func (b *destinatarioBuilder) SetEndereco(endereco string) *destinatarioBuilder {
	b.destinatario.endereco = endereco
	return b
}

func (b *destinatarioBuilder) Build() (Destinatario, error) {
	v := validator.New()

	v.CheckBlank("nome do destinatário", b.destinatario.nome)
	v.CheckBlank("endereco do destinatário", b.destinatario.endereco)

	if v.HasErrors() {
		return nil, errs.NewValidationError(v.Errors)
	}

	return b.destinatario, nil
}
