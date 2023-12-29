package cliente

import (
	"go/mydelivery/model"
	"go/mydelivery/shared/errs"
)

type catalogoService struct {
	clientes model.ClienteRepository
}

func NewCatalogoService(clientes model.ClienteRepository) CatalogoService {
	return &catalogoService{clientes}
}

func (s *catalogoService) Salvar(input CatalogoClienteRequest) (CatalogoClienteResponse, error) {
	cliente, err := model.NewCliente().
		SetNome(input.Nome).
		SetEmail(input.Email).
		SetTelefone(input.Telefone).
		Build()

	if err != nil {
		return CatalogoClienteResponse{}, err
	}

	clienteExistente, err := s.clientes.ObterPorEmail(cliente.Email())
	if err != nil {
		return CatalogoClienteResponse{}, err
	}

	if clienteExistente != nil {
		return CatalogoClienteResponse{}, errs.NewConflictError("já existe um cliente cadastrado com esse email")
	}

	novoClienteID, err := s.clientes.Salvar(cliente)
	if err != nil {
		return CatalogoClienteResponse{}, err
	}

	response := ToClienteResponse(cliente)
	response.ID = novoClienteID

	return response, nil
}

func (s *catalogoService) Atualizar(id int64, input CatalogoClienteRequest) (CatalogoClienteResponse, error) {
	cliente, err := model.NewCliente().
		SetID(id).
		SetNome(input.Nome).
		SetEmail(input.Email).
		SetTelefone(input.Telefone).
		Build()

	if err != nil {
		return CatalogoClienteResponse{}, err
	}

	clienteExistente, err := s.clientes.ObterPorEmail(cliente.Email())
	if err != nil {
		return CatalogoClienteResponse{}, err
	}

	if clienteExistente != nil && clienteExistente.ID() != cliente.ID() {
		return CatalogoClienteResponse{}, errs.NewConflictError("já existe um cliente cadastrado com esse email")
	}

	err = s.clientes.Atualizar(cliente)
	if err != nil {
		return CatalogoClienteResponse{}, err
	}

	return ToClienteResponse(cliente), nil
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
