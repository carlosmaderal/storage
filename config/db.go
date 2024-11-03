package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	// Monta a string de conexão com variáveis de ambiente
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASW"),
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	// Cria as tabelas, se não existirem
	createTables()
}

func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS vault_hashes (
			id BIGINT PRIMARY KEY AUTO_INCREMENT,
			hashcheck LONGTEXT NOT NULL,
			content LONGBLOB,
			lastcheck BIGINT NOT NULL,
			registered BIGINT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS vault_files (
			id BIGINT PRIMARY KEY AUTO_INCREMENT,
			uid VARCHAR(50),
			filepath LONGTEXT NOT NULL,
			hashcheck LONGTEXT NOT NULL,
			lastseen BIGINT NOT NULL,
			lastchange BIGINT NOT NULL,
			registered BIGINT NOT NULL,
			tags LONGTEXT
		)`,
	}
	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Println("Debug", os.Getenv("DB_USER"))
			log.Fatalf("Erro ao criar tabela: %v", err)
		}
	}
}
