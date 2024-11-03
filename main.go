package main

import (
	"log"
	// "os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"storage/config"
	"storage/routes"
)

func main() {
	// Carrega variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
        log.Println("Aviso: arquivo .env não encontrado. Usando variáveis de ambiente padrão.")
	}

	// Inicializa conexão com o banco de dados
	config.InitDB()

	// Cria o router
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Inicia o servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
