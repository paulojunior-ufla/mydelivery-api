package entrega

import (
	"go/mydelivery/model"
	"time"
)

type SolicitaEntregaService interface {
	Solicitar(SolicitaEntregaRequest) (SolicitaEntregaResponse, error)
}

type SolicitaEntregaRequest struct {
	ClienteID            int64   `json:"cliente_id"`
	Taxa                 float64 `json:"taxa"`
	NomeDestinatario     string  `json:"nome_destinatario"`
	EnderecoDestinatario string  `json:"endereco_destinatario"`
}

type SolicitaEntregaResponse struct {
	ID                   int64     `json:"id"`
	NomeCliente          string    `json:"cliente"`
	Taxa                 float64   `json:"taxa"`
	DataPedido           time.Time `json:"data_pedido"`
	NomeDestinatario     string    `json:"nome_destinatario"`
	EnderecoDestinatario string    `json:"endereco_destinatario"`
}

func ToSolicitaEntregaResponse(e model.Entrega) SolicitaEntregaResponse {
	return SolicitaEntregaResponse{
		ID:                   e.ID(),
		NomeCliente:          e.Cliente().Nome(),
		Taxa:                 e.Taxa(),
		DataPedido:           e.DataPedido(),
		NomeDestinatario:     e.Destinatario().Nome(),
		EnderecoDestinatario: e.Destinatario().Endereco(),
	}
}
