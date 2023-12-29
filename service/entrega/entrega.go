package entrega

import (
	"go/mydelivery/model"
	"time"
)

type SolicitaEntregaService interface {
	Solicitar(SolicitaEntregaRequest) (EntregaResponse, error)
}

type FinalizaEntregaService interface {
	Finalizar(idEntrega int64) error
}

type SolicitaEntregaRequest struct {
	ClienteID            int64   `json:"cliente_id"`
	Taxa                 float64 `json:"taxa"`
	NomeDestinatario     string  `json:"nome_destinatario"`
	EnderecoDestinatario string  `json:"endereco_destinatario"`
}

type EntregaResponse struct {
	ID                   int64     `json:"id"`
	ClienteID            int64     `json:"cliente_id"`
	Taxa                 float64   `json:"taxa"`
	Status               string    `json:"status"`
	DataPedido           time.Time `json:"data_pedido"`
	NomeDestinatario     string    `json:"nome_destinatario"`
	EnderecoDestinatario string    `json:"endereco_destinatario"`
}

func ToEntregaResponse(e model.Entrega) EntregaResponse {
	return EntregaResponse{
		ID:                   e.ID(),
		ClienteID:            e.Cliente().ID(),
		Taxa:                 e.Taxa(),
		Status:               string(e.Status()),
		DataPedido:           e.DataPedido(),
		NomeDestinatario:     e.Destinatario().Nome(),
		EnderecoDestinatario: e.Destinatario().Endereco(),
	}
}

func ToEntregaResponseCollection(ee []model.Entrega) []EntregaResponse {
	response := []EntregaResponse{}
	for _, e := range ee {
		response = append(response, ToEntregaResponse(e))
	}
	return response
}
