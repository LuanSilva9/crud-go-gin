package main

import "database/sql"

type Pessoa struct {
	ID   int
	Nome string `json:"nome"`
}

func criarPessoa(db *sql.DB, p Pessoa) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO pessoa(nome) VALUES ($1) RETURNING id", p.Nome).Scan(&id)
	return id, err
}

func lerPessoa(db *sql.DB, id int) (Pessoa, error) {
	var p Pessoa
	err := db.QueryRow("SELECT id, nome FROM pessoa where id = $1", id).Scan(&p.ID, &p.Nome)
	return p, err
}

func atualizarPessoa(db *sql.DB, id int, nome string) error {
	_, err := db.Exec("UPDATE pessoa SET nome=$1 WHERE id=$2", nome, id)
	return err
}

func deletarPessoa(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM pessoa WHERE id=$1", id)
	return err
}
