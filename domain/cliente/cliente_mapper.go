package cliente

func ToClienteResponse(c Cliente) ClienteResponse {
	return ClienteResponse{
		ID:       c.ID(),
		Nome:     c.Nome(),
		Email:    c.Email(),
		Telefone: c.Telefone(),
	}
}

func ToClienteResponseCollection(cc []Cliente) []ClienteResponse {
	response := []ClienteResponse{}
	for _, c := range cc {
		response = append(response, ToClienteResponse(c))
	}
	return response
}
