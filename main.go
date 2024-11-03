package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Marca o início do tempo total
    totalStartTime := time.Now()

    // Captura o valor da variável de ambiente ENV
    env := os.Getenv("ENV")
    fmt.Println("Valor de ENV:", env)

    // Configurações do banco de dados
    host := os.Getenv("DB_HOST")
    dbname := os.Getenv("DB_NAME")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASS")

    // Marca o início do tempo de carregamento do banco de dados
    dbStartTime := time.Now()

    // Cria a conexão
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("\nErro ao conectar ao banco de dados: %v", err)
    }
    defer db.Close()

    // Testa a conexão
    if err = db.Ping(); err != nil {
        log.Fatalf("\nErro ao conectar ao banco de dados: %v", err)
    }

    fmt.Println("\nConexão com o banco de dados realizada com sucesso!")

    // Calcula o tempo de carregamento do banco de dados
    dbLoadTime := time.Since(dbStartTime).Seconds()
    fmt.Printf("\nTempo de carregamento do banco de dados: %.4f segundos\n", dbLoadTime)

    // Calcula o tempo total de carregamento
    totalLoadTime := time.Since(totalStartTime).Seconds()
    fmt.Printf("Tempo total de carregamento: %.4f segundos\n", totalLoadTime)
}
