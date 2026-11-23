package servicos

import (
	"strings"

	"de-mala-pronta-api/daos"
	"de-mala-pronta-api/modelos"
)

type DestinoServico struct {
	dao       *daos.DestinoDAO
	viagemDAO *daos.ViagemDAO
}

func NovoDestinoServico(dao *daos.DestinoDAO, viagemDAO *daos.ViagemDAO) *DestinoServico {
	return &DestinoServico{dao: dao, viagemDAO: viagemDAO}
}

func (s *DestinoServico) validar(d *modelos.Destino) error {
	if d.ViagemID <= 0 || strings.TrimSpace(d.Cidade) == "" || strings.TrimSpace(d.Pais) == "" ||
		d.DataChegada == "" || d.DataSaida == "" {
		return erroInvalido("todos os campos são obrigatórios")
	}
	chegada, err := parseData(d.DataChegada)
	if err != nil {
		return erroInvalido("data_chegada deve estar no formato YYYY-MM-DD")
	}
	saida, err := parseData(d.DataSaida)
	if err != nil {
		return erroInvalido("data_saida deve estar no formato YYYY-MM-DD")
	}
	if chegada.After(saida) {
		return erroInvalido("data_chegada não pode ser posterior a data_saida")
	}
	_, err = buscarViagemRelacionada(s.viagemDAO, d.ViagemID)
	return err
}

func (s *DestinoServico) Cadastrar(d *modelos.Destino) error {
	if err := s.validar(d); err != nil {
		return err
	}
	return s.dao.Inserir(d)
}

func (s *DestinoServico) Listar(viagemID int) ([]modelos.Destino, error) {
	return s.dao.Listar(viagemID)
}

func (s *DestinoServico) BuscarPorID(id int) (*modelos.Destino, error) {
	d, err := s.dao.BuscarPorID(id)
	if err != nil {
		return nil, converteErro(err)
	}
	return d, nil
}

func (s *DestinoServico) Atualizar(d *modelos.Destino) error {
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

func (s *DestinoServico) Excluir(id int) error {
	linhas, err := s.dao.Excluir(id)
	if err != nil {
		return converteErro(err)
	}
	if linhas == 0 {
		return ErrNaoEncontrado
	}
	return nil
}
