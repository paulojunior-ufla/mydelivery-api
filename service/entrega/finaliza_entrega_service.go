package entrega

import (
	"go/mydelivery/model"
	"go/mydelivery/shared/errs"
)

type finalizaEntregaService struct {
	entregas model.EntregaRepository
}

func NewFinalizaEntregaService(entregas model.EntregaRepository) FinalizaEntregaService {
	return &finalizaEntregaService{entregas}
}

func (s *finalizaEntregaService) Finalizar(id int64) (EntregaResponse, error) {
	entrega, err := s.entregas.ObterPorID(id)
	if err != nil {
		return EntregaResponse{}, err
	}

	if entrega == nil {
		return EntregaResponse{}, errs.NewNotFoundError("entrega não encontrada")
	}

	err = entrega.Finalizar()
	if err != nil {
		return EntregaResponse{}, err
	}

	return ToEntregaResponse(entrega), nil

}
