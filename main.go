package main

import (
	"log"
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

	// Middleware CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Registra rotas
	routes.RegisterRoutes(r)

	// Inicia o servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
