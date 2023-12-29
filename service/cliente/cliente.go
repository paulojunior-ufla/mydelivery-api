package cliente

import "go/mydelivery/model"

type CatalogoService interface {
	Salvar(CatalogoClienteRequest) (CatalogoClienteResponse, error)
	Atualizar(int64, CatalogoClienteRequest) (CatalogoClienteResponse, error)
}

type CatalogoClienteRequest struct {
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
}

type CatalogoClienteResponse struct {
	ID       int64  `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
}

func ToClienteResponse(c model.Cliente) CatalogoClienteResponse {
	return CatalogoClienteResponse{
		ID:       c.ID(),
		Nome:     c.Nome(),
		Email:    c.Email(),
		Telefone: c.Telefone(),
	}
}

func ToClienteResponseCollection(cc []model.Cliente) []CatalogoClienteResponse {
	response := []CatalogoClienteResponse{}
	for _, c := range cc {
		response = append(response, ToClienteResponse(c))
	}
	return response
}
