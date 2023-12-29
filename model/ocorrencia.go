package model

import (
	"go/mydelivery/shared/errs"
	"go/mydelivery/shared/validator"
	"time"
)

type OcorrenciaRepository interface {
	Salvar(o Ocorrencia) (int64, error)
}

type Ocorrencia interface {
	ID() int64
	Entrega() Entrega
	Descricao() string
	DataRegistro() time.Time
}

type ocorrencia struct {
	id           int64
	entrega      Entrega
	descricao    string
	dataRegistro time.Time
}

func (o *ocorrencia) ID() int64               { return o.id }
func (o *ocorrencia) Entrega() Entrega        { return o.entrega }
func (o *ocorrencia) Descricao() string       { return o.descricao }
func (o *ocorrencia) DataRegistro() time.Time { return o.dataRegistro }

type ocorrenciaBuilder struct {
	ocorrencia *ocorrencia
}

func NewOcorrencia() *ocorrenciaBuilder {
	return &ocorrenciaBuilder{
		ocorrencia: &ocorrencia{
			dataRegistro: time.Now(),
		},
	}
}

func (b *ocorrenciaBuilder) SetID(id int64) *ocorrenciaBuilder {
	b.ocorrencia.id = id
	return b
}

func (b *ocorrenciaBuilder) SetEntrega(e Entrega) *ocorrenciaBuilder {
	b.ocorrencia.entrega = e
	return b
}

func (b *ocorrenciaBuilder) SetDescricao(desc string) *ocorrenciaBuilder {
	b.ocorrencia.descricao = desc
	return b
}

func (b *ocorrenciaBuilder) Build() (Ocorrencia, error) {
	v := validator.New()

	v.Check(b.ocorrencia.entrega != nil, "entrega da ocorrência", "não pode ser vazio")
	v.CheckBlank("descrição da ocorrência", b.ocorrencia.descricao)

	if v.HasErrors() {
		return nil, errs.NewValidationError(v.Errors)
	}

	return b.ocorrencia, nil
}
