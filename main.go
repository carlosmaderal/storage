package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Carrega as variáveis de ambiente
	env := os.Getenv("ENV")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASW")

	// Monta a string de conexão
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	// Conecta ao banco de dados
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer db.Close()

	// Testa a conexão
	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	// Cria o router
	r := gin.Default()

	// Rota inicial
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, " V2 - Conexão bem-sucedida ao banco de dados em %s com ambiente %s", dbName, env)
	})

	// Inicia o servidor
	r.Run(":8080")
}
