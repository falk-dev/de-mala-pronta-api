package daos

import (
	"database/sql"

	"de-mala-pronta-api/modelos"
)

type DestinoDAO struct {
	db *sql.DB
}

func NovoDestinoDAO(db *sql.DB) *DestinoDAO {
	return &DestinoDAO{db: db}
}

func (d *DestinoDAO) Inserir(m *modelos.Destino) error {
	res, err := d.db.Exec(
		"INSERT INTO destinos (viagem_id, cidade, pais, data_chegada, data_saida) VALUES (?, ?, ?, ?, ?)",
		m.ViagemID, m.Cidade, m.Pais, m.DataChegada, m.DataSaida)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	m.ID = int(id)
	return err
}

func (d *DestinoDAO) Listar(viagemID int) ([]modelos.Destino, error) {
	consulta := "SELECT id, viagem_id, cidade, pais, data_chegada, data_saida FROM destinos"
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

	destinos := []modelos.Destino{}
	for linhas.Next() {
		var m modelos.Destino
		if err := linhas.Scan(&m.ID, &m.ViagemID, &m.Cidade, &m.Pais, &m.DataChegada, &m.DataSaida); err != nil {
			return nil, err
		}
		destinos = append(destinos, m)
	}
	return destinos, linhas.Err()
}

func (d *DestinoDAO) BuscarPorID(id int) (*modelos.Destino, error) {
	var m modelos.Destino
	err := d.db.QueryRow(
		"SELECT id, viagem_id, cidade, pais, data_chegada, data_saida FROM destinos WHERE id = ?", id).
		Scan(&m.ID, &m.ViagemID, &m.Cidade, &m.Pais, &m.DataChegada, &m.DataSaida)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (d *DestinoDAO) Atualizar(m *modelos.Destino) (int64, error) {
	res, err := d.db.Exec(
		"UPDATE destinos SET viagem_id = ?, cidade = ?, pais = ?, data_chegada = ?, data_saida = ? WHERE id = ?",
		m.ViagemID, m.Cidade, m.Pais, m.DataChegada, m.DataSaida, m.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *DestinoDAO) Excluir(id int) (int64, error) {
	res, err := d.db.Exec("DELETE FROM destinos WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
