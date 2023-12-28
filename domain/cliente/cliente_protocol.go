package cliente

type Cliente interface {
	ID() int64
	Nome() string
	Email() string
	Telefone() string
}

type Repository interface {
	Todos() ([]Cliente, error)
	ObterPorID(int64) (Cliente, error)
	ObterPorEmail(string) (Cliente, error)
	Salvar(Cliente) (int64, error)
	Atualizar(cliente Cliente) error
	Excluir(id int64) error
}

type Service interface {
	Salvar(ClienteRequest) (ClienteResponse, error)
	Atualizar(int64, ClienteRequest) (ClienteResponse, error)
}

type ClienteRequest struct {
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
}

type ClienteResponse struct {
	ID       int64  `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
}

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
