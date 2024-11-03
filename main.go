package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
    env := os.Getenv("ENV")
    fmt.Fprintf(w, "Valor de ENV: %s\n", env)

    host := os.Getenv("DB_HOST")
    dbname := os.Getenv("DB_NAME")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASS")

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    if err = db.Ping(); err != nil {
        http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Conexão com o banco de dados realizada com sucesso!\n")
}

func main() {
    http.HandleFunc("/", handler)

    log.Fatal(http.ListenAndServe(":8000", nil))
}
