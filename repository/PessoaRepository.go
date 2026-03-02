package repository

import (
	"database/sql"

	"github.com/LuanSilva9/crud-go-gin/model"
)

func CriarPessoa(db *sql.DB, p model.Pessoa) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO pessoa(nome) VALUES ($1) RETURNING id", p.Nome).Scan(&id)
	return id, err
}

func LerPessoa(db *sql.DB, id int) (model.Pessoa, error) {
	var p model.Pessoa
	err := db.QueryRow("SELECT id, nome FROM pessoa where id = $1", id).Scan(&p.ID, &p.Nome)
	return p, err
}

func AtualizarPessoa(db *sql.DB, id int, nome string) error {
	_, err := db.Exec("UPDATE pessoa SET nome=$1 WHERE id=$2", nome, id)
	return err
}

func DeletarPessoa(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM pessoa WHERE id=$1", id)
	return err
}
