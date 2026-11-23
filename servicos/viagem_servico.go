package servicos

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"de-mala-pronta-api/daos"
	"de-mala-pronta-api/modelos"
)

var (
	ErrInvalido       = errors.New("dados inválidos")
	ErrNaoEncontrado  = errors.New("registro não encontrado")
	ErrConflito       = errors.New("conflito de relacionamento")
	formatoDataPadrao = "2006-01-02"
)

func erroInvalido(msg string) error {
	return fmt.Errorf("%w: %s", ErrInvalido, msg)
}

func buscarViagemRelacionada(dao *daos.ViagemDAO, id int) (*modelos.Viagem, error) {
	v, err := dao.BuscarPorID(id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w: a viagem informada não existe", ErrConflito)
	}
	return v, err
}

func converteErro(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNaoEncontrado
	}
	if strings.Contains(err.Error(), "FOREIGN KEY") {
		return ErrConflito
	}
	return err
}

func parseData(texto string) (time.Time, error) {
	return time.Parse(formatoDataPadrao, texto)
}

type ViagemServico struct {
	dao *daos.ViagemDAO
}

func NovoViagemServico(dao *daos.ViagemDAO) *ViagemServico {
	return &ViagemServico{dao: dao}
}

func (s *ViagemServico) validar(v *modelos.Viagem) error {
	if strings.TrimSpace(v.Titulo) == "" || strings.TrimSpace(v.Status) == "" ||
		v.DataInicio == "" || v.DataFim == "" {
		return erroInvalido("todos os campos são obrigatórios")
	}
	inicio, err := parseData(v.DataInicio)
	if err != nil {
		return erroInvalido("data_inicio deve estar no formato YYYY-MM-DD")
	}
	fim, err := parseData(v.DataFim)
	if err != nil {
		return erroInvalido("data_fim deve estar no formato YYYY-MM-DD")
	}
	if fim.Before(inicio) {
		return erroInvalido("data_fim não pode ser anterior a data_inicio")
	}
	if v.Orcamento < 0 {
		return erroInvalido("orcamento não pode ser negativo")
	}
	return nil
}

func (s *ViagemServico) Cadastrar(v *modelos.Viagem) error {
	if err := s.validar(v); err != nil {
		return err
	}
	return s.dao.Inserir(v)
}

func (s *ViagemServico) Listar() ([]modelos.Viagem, error) {
	return s.dao.Listar()
}

func (s *ViagemServico) BuscarPorID(id int) (*modelos.Viagem, error) {
	v, err := s.dao.BuscarPorID(id)
	if err != nil {
		return nil, converteErro(err)
	}
	return v, nil
}

func (s *ViagemServico) Atualizar(v *modelos.Viagem) error {
	if err := s.validar(v); err != nil {
		return err
	}
	linhas, err := s.dao.Atualizar(v)
	if err != nil {
		return converteErro(err)
	}
	if linhas == 0 {
		return ErrNaoEncontrado
	}
	return nil
}

func (s *ViagemServico) Excluir(id int) error {
	linhas, err := s.dao.Excluir(id)
	if err != nil {
		return converteErro(err)
	}
	if linhas == 0 {
		return ErrNaoEncontrado
	}
	return nil
}
