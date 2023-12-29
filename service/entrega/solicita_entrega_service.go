package entrega

import (
	"go/mydelivery/model"
	"go/mydelivery/shared/errs"
)

type solicitaEntregaService struct {
	clientes model.ClienteRepository
	entregas model.EntregaRepository
}

func NewSolicitaEntregaService(clientes model.ClienteRepository, entregas model.EntregaRepository) SolicitaEntregaService {
	return &solicitaEntregaService{clientes, entregas}
}

func (s *solicitaEntregaService) Solicitar(input SolicitaEntregaRequest) (SolicitaEntregaResponse, error) {
	cliente, err := s.clientes.ObterPorID(input.ClienteID)
	if err != nil {
		return SolicitaEntregaResponse{}, err
	}

	if cliente == nil {
		return SolicitaEntregaResponse{}, errs.NewConflictError("cliente inválido")
	}

	entrega, err := model.NewEntrega().
		SetCliente(cliente).
		SetTaxa(input.Taxa).
		Build()

	if err != nil {
		return SolicitaEntregaResponse{}, err
	}

	idNovaEntrega, err := s.entregas.Salvar(entrega)
	if err != nil {
		return SolicitaEntregaResponse{}, err
	}

	return SolicitaEntregaResponse{ID: idNovaEntrega}, nil
}
