package service

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/LuanSilva9/crud-go-gin/model"
	"github.com/LuanSilva9/crud-go-gin/repository"
	"github.com/gin-gonic/gin"
)

type PessoaService struct {
	DB *sql.DB
}

func (s *PessoaService) CriarPessoaHandler(c *gin.Context) {
	var p model.Pessoa

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := repository.CriarPessoa(s.DB, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	p.ID = id

	c.JSON(http.StatusCreated, p)
}

func (s *PessoaService) LerPessoaHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p, err := repository.LerPessoa(s.DB, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pessoa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, p)
}

func (s *PessoaService) AtualizarPessoaHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var p model.Pessoa

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.AtualizarPessoa(s.DB, id, p.Nome); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "atualizado"})
}

func (s *PessoaService) DeletarPessoaHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := repository.DeletarPessoa(s.DB, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deletado"})
}
