package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    totalStartTime := time.Now()

    env := os.Getenv("ENV")
    fmt.Println("Valor de ENV:", env)

    host := os.Getenv("DB_HOST")
    dbname := os.Getenv("DB_NAME")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASS")

    dbStartTime := time.Now()

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("\nErro ao conectar ao banco de dados: %v", err)
    }
    defer db.Close()

    if err = db.Ping(); err != nil {
        log.Fatalf("\nErro ao conectar ao banco de dados: %v", err)
    }

    fmt.Println("\nConexão com o banco de dados realizada com sucesso!")

    dbLoadTime := time.Since(dbStartTime).Seconds()
    fmt.Printf("\nTempo de carregamento do banco de dados: %.4f segundos\n", dbLoadTime)

    totalLoadTime := time.Since(totalStartTime).Seconds()
    fmt.Printf("Tempo total de carregamento: %.4f segundos\n", totalLoadTime)

    // Inicia o servidor na porta definida pelo Cloud Run
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Valor padrão
    }
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
