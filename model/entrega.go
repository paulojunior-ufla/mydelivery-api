package model

import (
	"go/mydelivery/shared/errs"
	"go/mydelivery/shared/validator"
	"time"
)

type EntregaRepository interface {
	Todos() ([]Entrega, error)
	ObterPorID(int64) (Entrega, error)
	Salvar(Entrega) (int64, error)
	Atualizar(Entrega) error
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
	DataPedido() time.Time
	Destinatario() Destinatario

	Finalizar() error
}

type entrega struct {
	id           int64
	cliente      Cliente
	taxa         float64
	status       StatusEntrega
	dataPedido   time.Time
	destinatario Destinatario
}

func (e *entrega) ID() int64                  { return e.id }
func (e *entrega) Cliente() Cliente           { return e.cliente }
func (e *entrega) Taxa() float64              { return e.taxa }
func (e *entrega) Status() StatusEntrega      { return e.status }
func (e *entrega) DataPedido() time.Time      { return e.dataPedido }
func (e *entrega) Destinatario() Destinatario { return e.destinatario }

func (e *entrega) podeSerFinalizada() bool    { return e.status == ENTREGA_PENDENTE }
func (e *entrega) naoPodeSerFinalizada() bool { return !e.podeSerFinalizada() }

func (e *entrega) Finalizar() error {
	if e.naoPodeSerFinalizada() {
		return errs.NewConflictError("esta entrega não pode ser finalizada")
	}
	e.status = ENTREGA_FINALIZADA
	return nil
}

type entregaBuilder struct {
	entrega *entrega
}

func NewEntrega() *entregaBuilder {
	return &entregaBuilder{
		entrega: &entrega{
			status:     ENTREGA_PENDENTE,
			dataPedido: time.Now(),
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

func (b *entregaBuilder) SetDestinatario(d Destinatario) *entregaBuilder {
	b.entrega.destinatario = d
	return b
}

func (b *entregaBuilder) Build() (Entrega, error) {
	v := validator.New()

	v.Check(b.entrega.cliente != nil, "cliente da entrega", "não pode ser vazio")
	v.Check(b.entrega.taxa >= 0, "taxa da entrega", "não pode ser negativa")
	v.Check(b.entrega.destinatario != nil, "destinatário da entrega", "não pode ser vazio")

	if v.HasErrors() {
		return nil, errs.NewValidationError(v.Errors)
	}

	return b.entrega, nil
}
