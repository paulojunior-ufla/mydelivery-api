package ocorrencia

import (
	"go/mydelivery/model"
	"go/mydelivery/shared/errs"
)

type registraOcorrenciaService struct {
	ocorrencias model.OcorrenciaRepository
	entregas    model.EntregaRepository
}

func NewRegistraOcorrenciaService(ocorrencias model.OcorrenciaRepository, entregas model.EntregaRepository) RegistraOcorrenciaService {
	return &registraOcorrenciaService{ocorrencias, entregas}
}

func (s *registraOcorrenciaService) Registrar(idEntrega int64, input RegistraOcorrenciaRequest) (RegistraOcorrenciaResponse, error) {
	entrega, err := s.entregas.ObterPorID(idEntrega)
	if err != nil {
		return RegistraOcorrenciaResponse{}, err
	}

	if entrega == nil {
		return RegistraOcorrenciaResponse{}, errs.NewNotFoundError("entrega não encontrada")
	}

	ocorrencia, err := model.NewOcorrencia().
		SetEntrega(entrega).
		SetDescricao(input.Descricao).
		Build()

	if err != nil {
		return RegistraOcorrenciaResponse{}, err
	}

	idNovaOcorrencia, err := s.ocorrencias.Salvar(ocorrencia)
	if err != nil {
		return RegistraOcorrenciaResponse{}, err
	}

	response := ToRegistraOcorrenciaResponse(ocorrencia)
	response.ID = idNovaOcorrencia

	return response, nil
}
