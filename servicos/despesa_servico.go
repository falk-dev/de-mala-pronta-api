package servicos

import (
	"strings"

	"de-mala-pronta-api/daos"
	"de-mala-pronta-api/modelos"
)

type DespesaServico struct {
	dao       *daos.DespesaDAO
	viagemDAO *daos.ViagemDAO
}

func NovoDespesaServico(dao *daos.DespesaDAO, viagemDAO *daos.ViagemDAO) *DespesaServico {
	return &DespesaServico{dao: dao, viagemDAO: viagemDAO}
}

func (s *DespesaServico) validar(d *modelos.Despesa) error {
	if d.ViagemID <= 0 || strings.TrimSpace(d.Descricao) == "" || strings.TrimSpace(d.Categoria) == "" ||
		d.Data == "" {
		return erroInvalido("todos os campos são obrigatórios")
	}
	if d.Valor <= 0 {
		return erroInvalido("valor deve ser maior que zero")
	}
	data, err := parseData(d.Data)
	if err != nil {
		return erroInvalido("data deve estar no formato YYYY-MM-DD")
	}
	viagem, err := buscarViagemRelacionada(s.viagemDAO, d.ViagemID)
	if err != nil {
		return err
	}
	inicio, _ := parseData(viagem.DataInicio)
	fim, _ := parseData(viagem.DataFim)
	if data.Before(inicio) || data.After(fim) {
		return erroInvalido("data da despesa deve estar dentro do período da viagem")
	}
	return nil
}

func (s *DespesaServico) Cadastrar(d *modelos.Despesa) error {
	if err := s.validar(d); err != nil {
		return err
	}
	return s.dao.Inserir(d)
}

func (s *DespesaServico) Listar(viagemID int) ([]modelos.Despesa, error) {
	return s.dao.Listar(viagemID)
}

func (s *DespesaServico) BuscarPorID(id int) (*modelos.Despesa, error) {
	d, err := s.dao.BuscarPorID(id)
	if err != nil {
		return nil, converteErro(err)
	}
	return d, nil
}

func (s *DespesaServico) Atualizar(d *modelos.Despesa) error {
	if err := s.validar(d); err != nil {
		return err
	}
	linhas, err := s.dao.Atualizar(d)
	if err != nil {
		return converteErro(err)
	}
	if linhas == 0 {
		return ErrNaoEncontrado
	}
	return nil
}

func (s *DespesaServico) Excluir(id int) error {
	linhas, err := s.dao.Excluir(id)
	if err != nil {
		return converteErro(err)
	}
	if linhas == 0 {
		return ErrNaoEncontrado
	}
	return nil
}

func (s *DespesaServico) TotalPorViagem(viagemID int) (float64, error) {
	if _, err := s.viagemDAO.BuscarPorID(viagemID); err != nil {
		return 0, converteErro(err)
	}
	return s.dao.TotalPorViagem(viagemID)
}
