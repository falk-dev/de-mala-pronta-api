package daos

import (
	"database/sql"

	"de-mala-pronta-api/modelos"
)

type ViagemDAO struct {
	db *sql.DB
}

func NovoViagemDAO(db *sql.DB) *ViagemDAO {
	return &ViagemDAO{db: db}
}

func (d *ViagemDAO) Inserir(v *modelos.Viagem) error {
	res, err := d.db.Exec(
		"INSERT INTO viagens (titulo, data_inicio, data_fim, orcamento, status) VALUES (?, ?, ?, ?, ?)",
		v.Titulo, v.DataInicio, v.DataFim, v.Orcamento, v.Status)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	v.ID = int(id)
	return err
}

func (d *ViagemDAO) Listar() ([]modelos.Viagem, error) {
	linhas, err := d.db.Query("SELECT id, titulo, data_inicio, data_fim, orcamento, status FROM viagens ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	viagens := []modelos.Viagem{}
	for linhas.Next() {
		var v modelos.Viagem
		if err := linhas.Scan(&v.ID, &v.Titulo, &v.DataInicio, &v.DataFim, &v.Orcamento, &v.Status); err != nil {
			return nil, err
		}
		viagens = append(viagens, v)
	}
	return viagens, linhas.Err()
}

func (d *ViagemDAO) BuscarPorID(id int) (*modelos.Viagem, error) {
	var v modelos.Viagem
	err := d.db.QueryRow(
		"SELECT id, titulo, data_inicio, data_fim, orcamento, status FROM viagens WHERE id = ?", id).
		Scan(&v.ID, &v.Titulo, &v.DataInicio, &v.DataFim, &v.Orcamento, &v.Status)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (d *ViagemDAO) Atualizar(v *modelos.Viagem) (int64, error) {
	res, err := d.db.Exec(
		"UPDATE viagens SET titulo = ?, data_inicio = ?, data_fim = ?, orcamento = ?, status = ? WHERE id = ?",
		v.Titulo, v.DataInicio, v.DataFim, v.Orcamento, v.Status, v.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *ViagemDAO) Excluir(id int) (int64, error) {
	res, err := d.db.Exec("DELETE FROM viagens WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
