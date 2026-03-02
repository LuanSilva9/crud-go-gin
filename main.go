package main

import (
	"database/sql"
	"fmt"

	"github.com/LuanSilva9/crud-go-gin/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error

	var connStr = connect()

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Conectado ao Banco")

	r := gin.Default()

	pessoaService := service.PessoaService{DB: db}

	r.POST("/pessoa", pessoaService.CriarPessoaHandler)
	r.GET("/pessoa/:id", pessoaService.LerPessoaHandler)
	r.PUT("/pessoa/:id", pessoaService.AtualizarPessoaHandler)
	r.DELETE("/pessoa/:id", pessoaService.DeletarPessoaHandler)

	r.Run(":8080")
}
