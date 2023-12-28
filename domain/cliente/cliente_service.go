package cliente

import "go/mydelivery/shared/errs"

type salvaClienteService struct {
	clientes Repository
}

func NewSalvaClienteService(cc Repository) ClienteService {
	return &salvaClienteService{cc}
}

func (s *salvaClienteService) Salvar(input ClienteRequest) (ClienteResponse, error) {
	cliente := New(input.Nome, input.Email, input.Telefone)
	err := cliente.Validar()
	if err != nil {
		return ClienteResponse{}, err
	}

	clienteExistente, err := s.clientes.ObterPorEmail(input.Email)
	if err != nil {
		return ClienteResponse{}, err
	}

	if clienteExistente != nil {
		return ClienteResponse{}, errs.NewConflictError("já existe um cliente cadastrado com esse email")
	}

	novoCliente, err := s.clientes.Salvar(cliente)
	if err != nil {
		return ClienteResponse{}, err
	}

	return ClienteResponse{
		ID:       novoCliente.ID(),
		Nome:     novoCliente.Nome(),
		Email:    novoCliente.Email(),
		Telefone: novoCliente.Telefone(),
	}, nil
}

func (s *salvaClienteService) Atualizar(id int64, input ClienteRequest) (ClienteResponse, error) {
	cliente := NewWithID(id, input.Nome, input.Email, input.Telefone)
	err := cliente.Validar()
	if err != nil {
		return ClienteResponse{}, err
	}

	clienteExistente, err := s.clientes.ObterPorEmail(input.Email)
	if err != nil {
		return ClienteResponse{}, err
	}

	if clienteExistente != nil && clienteExistente.ID() != cliente.ID() {
		return ClienteResponse{}, errs.NewConflictError("já existe um cliente cadastrado com esse email")
	}

	clienteAtualizado, err := s.clientes.Atualizar(cliente)
	if err != nil {
		return ClienteResponse{}, err
	}

	return ClienteResponse{
		ID:       clienteAtualizado.ID(),
		Nome:     clienteAtualizado.Nome(),
		Email:    clienteAtualizado.Email(),
		Telefone: clienteAtualizado.Telefone(),
	}, nil
}
