package ocorrencia

import (
	"go/mydelivery/model"
	"time"
)

type RegistraOcorrenciaService interface {
	Registrar(idEntrega int64, input RegistraOcorrenciaRequest) (RegistraOcorrenciaResponse, error)
}

type RegistraOcorrenciaRequest struct {
	Descricao string `json:"descricao"`
}

type RegistraOcorrenciaResponse struct {
	ID           int64     `json:"id"`
	EntregaID    int64     `json:"entrega_id"`
	Descricao    string    `json:"descricao"`
	DataRegistro time.Time `json:"data_registro"`
}

func ToRegistraOcorrenciaResponse(o model.Ocorrencia) RegistraOcorrenciaResponse {
	return RegistraOcorrenciaResponse{
		ID:           o.ID(),
		EntregaID:    o.Entrega().ID(),
		Descricao:    o.Descricao(),
		DataRegistro: o.DataRegistro(),
	}
}
