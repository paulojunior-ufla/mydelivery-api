package cliente

import "go/mydelivery/shared/errs"

type service struct {
	clientes Repository
}

func NewService(clientes Repository) Service {
	return &service{clientes}
}

func (s *service) Salvar(input ClienteRequest) (ClienteResponse, error) {
	cliente, err := New().
		SetNome(input.Nome).
		SetEmail(input.Email).
		SetTelefone(input.Telefone).
		Build()

	if err != nil {
		return ClienteResponse{}, err
	}

	clienteExistente, err := s.clientes.ObterPorEmail(cliente.Email())
	if err != nil {
		return ClienteResponse{}, err
	}

	if clienteExistente != nil {
		return ClienteResponse{}, errs.NewConflictError("já existe um cliente cadastrado com esse email")
	}

	id, err := s.clientes.Salvar(cliente)
	if err != nil {
		return ClienteResponse{}, err
	}

	response := ToClienteResponse(cliente)
	response.ID = id

	return response, nil
}

func (s *service) Atualizar(id int64, input ClienteRequest) (ClienteResponse, error) {
	cliente, err := New().
		SetID(id).
		SetNome(input.Nome).
		SetEmail(input.Email).
		SetTelefone(input.Telefone).
		Build()

	if err != nil {
		return ClienteResponse{}, err
	}

	clienteExistente, err := s.clientes.ObterPorEmail(cliente.Email())
	if err != nil {
		return ClienteResponse{}, err
	}

	if clienteExistente != nil && clienteExistente.ID() != cliente.ID() {
		return ClienteResponse{}, errs.NewConflictError("já existe um cliente cadastrado com esse email")
	}

	err = s.clientes.Atualizar(cliente)
	if err != nil {
		return ClienteResponse{}, err
	}

	return ToClienteResponse(cliente), nil
}
