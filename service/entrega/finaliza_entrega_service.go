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

func (s *finalizaEntregaService) Finalizar(id int64) error {
	entrega, err := s.entregas.ObterPorID(id)
	if err != nil {
		return err
	}

	if entrega == nil {
		return errs.NewNotFoundError("entrega não encontrada")
	}

	err = entrega.Finalizar()
	if err != nil {
		return err
	}

	err = s.entregas.Atualizar(entrega)
	if err != nil {
		return err
	}

	return nil
}
