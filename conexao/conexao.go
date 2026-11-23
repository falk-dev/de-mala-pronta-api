package conexao

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Abrir(caminho string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", caminho)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, err
	}
	if err := criarTabelas(db); err != nil {
		return nil, err
	}
	return db, nil
}

func criarTabelas(db *sql.DB) error {
	comandos := []string{
		`CREATE TABLE IF NOT EXISTS viagens (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			titulo TEXT NOT NULL,
			data_inicio TEXT NOT NULL,
			data_fim TEXT NOT NULL,
			orcamento REAL NOT NULL,
			status TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS destinos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			viagem_id INTEGER NOT NULL,
			cidade TEXT NOT NULL,
			pais TEXT NOT NULL,
			data_chegada TEXT NOT NULL,
			data_saida TEXT NOT NULL,
			FOREIGN KEY (viagem_id) REFERENCES viagens(id)
		)`,
		`CREATE TABLE IF NOT EXISTS despesas (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			viagem_id INTEGER NOT NULL,
			descricao TEXT NOT NULL,
			valor REAL NOT NULL,
			data TEXT NOT NULL,
			categoria TEXT NOT NULL,
			FOREIGN KEY (viagem_id) REFERENCES viagens(id)
		)`,
	}
	for _, c := range comandos {
		if _, err := db.Exec(c); err != nil {
			return err
		}
	}
	return nil
}
