package entrega

type SolicitaEntregaService interface {
	Solicitar(SolicitaEntregaRequest) (SolicitaEntregaResponse, error)
}

type SolicitaEntregaRequest struct {
	ClienteID int64   `json:"cliente_id"`
	Taxa      float64 `json:"taxa"`
}

type SolicitaEntregaResponse struct {
	ID int64 `json:"id"`
}
