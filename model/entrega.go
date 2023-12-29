package model

import (
	"go/mydelivery/shared/errs"
	"go/mydelivery/shared/validator"
)

type EntregaRepository interface {
	Salvar(Entrega) (int64, error)
}

type StatusEntrega string

const (
	ENTREGA_PENDENTE   StatusEntrega = "PENDENTE"
	ENTREGA_FINALIZADA StatusEntrega = "FINALIZADA"
	ENTREGA_CANCELADA  StatusEntrega = "CANCELADA"
)

type Entrega interface {
	ID() int64
	Cliente() Cliente
	Taxa() float64
	Status() StatusEntrega
}

type entrega struct {
	id      int64
	cliente Cliente
	taxa    float64
	status  StatusEntrega
}

func (e *entrega) ID() int64             { return e.id }
func (e *entrega) Cliente() Cliente      { return e.cliente }
func (e *entrega) Taxa() float64         { return e.taxa }
func (e *entrega) Status() StatusEntrega { return e.status }

type entregaBuilder struct {
	entrega *entrega
}

func NewEntrega() *entregaBuilder {
	return &entregaBuilder{
		entrega: &entrega{
			status: ENTREGA_PENDENTE,
		},
	}
}

func (b *entregaBuilder) SetID(id int64) *entregaBuilder {
	b.entrega.id = id
	return b
}

func (b *entregaBuilder) SetCliente(c Cliente) *entregaBuilder {
	b.entrega.cliente = c
	return b
}

func (b *entregaBuilder) SetTaxa(taxa float64) *entregaBuilder {
	b.entrega.taxa = taxa
	return b
}

func (b *entregaBuilder) SetStatus(status StatusEntrega) *entregaBuilder {
	b.entrega.status = status
	return b
}

func (b *entregaBuilder) Build() (Entrega, error) {
	v := validator.New()

	v.Check(b.entrega.cliente != nil, "cliente", "não pode ser vazio")
	v.Check(b.entrega.taxa >= 0, "taxa", "não pode ser negativa")

	statusEntregaValido := b.entrega.status == ENTREGA_PENDENTE || b.entrega.status == ENTREGA_FINALIZADA || b.entrega.status == ENTREGA_CANCELADA
	v.Check(statusEntregaValido, "status", "não é válido")

	if v.HasErrors() {
		return nil, errs.NewValidationError(v.Errors)
	}

	return b.entrega, nil
}