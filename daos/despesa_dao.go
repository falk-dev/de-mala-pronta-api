package daos

import (
	"database/sql"

	"de-mala-pronta-api/modelos"
)

type DespesaDAO struct {
	db *sql.DB
}

func NovoDespesaDAO(db *sql.DB) *DespesaDAO {
	return &DespesaDAO{db: db}
}

func (d *DespesaDAO) Inserir(m *modelos.Despesa) error {
	res, err := d.db.Exec(
		"INSERT INTO despesas (viagem_id, descricao, valor, data, categoria) VALUES (?, ?, ?, ?, ?)",
		m.ViagemID, m.Descricao, m.Valor, m.Data, m.Categoria)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	m.ID = int(id)
	return err
}

func (d *DespesaDAO) Listar(viagemID int) ([]modelos.Despesa, error) {
	consulta := "SELECT id, viagem_id, descricao, valor, data, categoria FROM despesas"
	args := []any{}
	if viagemID > 0 {
		consulta += " WHERE viagem_id = ?"
		args = append(args, viagemID)
	}
	consulta += " ORDER BY id"

	linhas, err := d.db.Query(consulta, args...)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	despesas := []modelos.Despesa{}
	for linhas.Next() {
		var m modelos.Despesa
		if err := linhas.Scan(&m.ID, &m.ViagemID, &m.Descricao, &m.Valor, &m.Data, &m.Categoria); err != nil {
			return nil, err
		}
		despesas = append(despesas, m)
	}
	return despesas, linhas.Err()
}

func (d *DespesaDAO) BuscarPorID(id int) (*modelos.Despesa, error) {
	var m modelos.Despesa
	err := d.db.QueryRow(
		"SELECT id, viagem_id, descricao, valor, data, categoria FROM despesas WHERE id = ?", id).
		Scan(&m.ID, &m.ViagemID, &m.Descricao, &m.Valor, &m.Data, &m.Categoria)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (d *DespesaDAO) Atualizar(m *modelos.Despesa) (int64, error) {
	res, err := d.db.Exec(
		"UPDATE despesas SET viagem_id = ?, descricao = ?, valor = ?, data = ?, categoria = ? WHERE id = ?",
		m.ViagemID, m.Descricao, m.Valor, m.Data, m.Categoria, m.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *DespesaDAO) Excluir(id int) (int64, error) {
	res, err := d.db.Exec("DELETE FROM despesas WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *DespesaDAO) TotalPorViagem(viagemID int) (float64, error) {
	var total float64
	err := d.db.QueryRow("SELECT COALESCE(SUM(valor), 0) FROM despesas WHERE viagem_id = ?", viagemID).Scan(&total)
	return total, err
}
