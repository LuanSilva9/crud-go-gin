# Golang CRUD API – Pessoa

API simples em Go para gerenciar pessoas, conectando ao **PostgreSQL**. Suporta operações **CRUD**: criar, ler, atualizar e deletar pessoas.

---

## Tecnologias

* [Golang](https://go.dev/)
* [Gin](https://github.com/gin-gonic/gin) – framework HTTP leve
* [PostgreSQL](https://www.postgresql.org/)
* [lib/pq](https://github.com/lib/pq) – driver PostgreSQL
* [godotenv](https://github.com/joho/godotenv) – leitura de variáveis de ambiente

---

## Estrutura do projeto

```
crud-go-gin/
 ├─ main.go                 # Bootstrap da aplicação e rotas
 ├─ pgconnect.go            # Conexão com PostgreSQL via .env
 ├─ model/
 │   └─ pessoa.go           # Entidade Pessoa (domain)
 ├─ repository/
 │   └─ pessoa_repository.go # Acesso ao banco (SQL)
 ├─ service/
 │   └─ pessoa_service.go   # Regras de negócio + handlers HTTP
 ├─ .env                    # Variáveis de ambiente
 ├─ go.mod
 └─ README.md
```

---

## Pré-requisitos

* Go >= 1.21
* PostgreSQL instalado e rodando
* Banco de dados criado (exemplo: `golang_crud`)

```sql
CREATE DATABASE golang_crud;

\c golang_crud

CREATE TABLE pessoa (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL
);
```

---

## Configuração do ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
DB_USER=root
DB_PASSWORD=1111
DB_HOST=localhost
DB_PORT=5432
DB_NAME=golang_crud
```

> **Não comite o `.env` com senha real em repositórios públicos.**

---

## Instalação

1. Clone o projeto:

```bash
git clone https://github.com/LuanSilva9/crud-go-gin.git
cd crud-go-gin
```

2. Inicialize as dependências:

```bash
go mod tidy
```

3. Instale Gin e godotenv:

```bash
go get github.com/gin-gonic/gin
"go get github.com/joho/godotenv"
go get github.com/lib/pq
```

---

## Rodando a API

Modo desenvolvimento:

```bash
go run .
```

* A API estará rodando em: `http://localhost:8080`

Modo produção (binário):

```bash
go build -o meuapp
./meuapp
```

---

## Endpoints

| Método | Endpoint      | Descrição                | Body (JSON)                |
| ------ | ------------- | ------------------------ | -------------------------- |
| POST   | `/pessoa`     | Criar nova pessoa        | `{ "nome": "Luan" }`       |
| GET    | `/pessoa/:id` | Ler pessoa por ID        | —                          |
| PUT    | `/pessoa/:id` | Atualizar nome da pessoa | `{ "nome": "Luan Alves" }` |
| DELETE | `/pessoa/:id` | Deletar pessoa por ID    | —                          |

---

## Observações

* O campo `ID` é gerado automaticamente pelo banco (auto-increment).
* Cada requisição HTTP já é processada em **goroutine**, garantindo concorrência nativa.
* Configure o pool de conexões no `main.go`:

```go
db.SetMaxOpenConns(20)
db.SetMaxIdleConns(10)
```

* Para produção, sempre use **.env** ou variáveis de ambiente para credenciais.

---

## Autor

Luan Silva – [GitHub](https://github.com/LuanSilva9)
