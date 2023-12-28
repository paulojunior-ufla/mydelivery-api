package cliente

type Cliente interface {
	ID() int64
	Nome() string
	Email() string
	Telefone() string
	Validar() error
}

type Repository interface {
	Todos() ([]Cliente, error)
	ObterPorID(int64) (Cliente, error)
	ObterPorEmail(string) (Cliente, error)
	Salvar(Cliente) (Cliente, error)
	Atualizar(cliente Cliente) (Cliente, error)
	Excluir(id int64) error
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

type ClienteService interface {
	Salvar(ClienteRequest) (ClienteResponse, error)
	Atualizar(int64, ClienteRequest) (ClienteResponse, error)
}
