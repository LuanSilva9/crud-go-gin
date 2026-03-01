package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

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

	r.POST("/pessoa", criarPessoaHandler)
	r.GET("/pessoa/:id", lerPessoaHandler)
	r.PUT("/pessoa/:id", atualizarPessoaHandler)
	r.DELETE("/pessoa/:id", deletarPessoaHandler)

	r.Run(":8080")
}

func criarPessoaHandler(c *gin.Context) {
	var p Pessoa

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := criarPessoa(db, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	p.ID = id

	c.JSON(http.StatusCreated, p)
}

func lerPessoaHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p, err := lerPessoa(db, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pessoa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, p)
}

func atualizarPessoaHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var p Pessoa

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := atualizarPessoa(db, id, p.Nome); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "atualizado"})
}

func deletarPessoaHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := deletarPessoa(db, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deletado"})
}
