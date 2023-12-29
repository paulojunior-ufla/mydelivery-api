package cliente

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
