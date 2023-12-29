package entrega

import (
	"go/mydelivery/model"
	"go/mydelivery/service/cliente"
	"time"
)

type SolicitaEntregaService interface {
	Solicitar(SolicitaEntregaRequest) (SolicitaEntregaResponse, error)
}

type SolicitaEntregaRequest struct {
	ClienteID int64   `json:"cliente_id"`
	Taxa      float64 `json:"taxa"`
}

type SolicitaEntregaResponse struct {
	ID         int64                           `json:"id"`
	Cliente    cliente.CatalogoClienteResponse `json:"cliente"`
	Taxa       float64                         `json:"taxa"`
	DataPedido time.Time                       `json:"data_pedido"`
}

func ToSolicitaEntregaResponse(e model.Entrega) SolicitaEntregaResponse {
	return SolicitaEntregaResponse{
		ID:         e.ID(),
		Cliente:    cliente.ToClienteResponse(e.Cliente()),
		Taxa:       e.Taxa(),
		DataPedido: e.DataPedido(),
	}
}
