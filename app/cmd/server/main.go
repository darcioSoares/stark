package main

import (
	"fmt"
	"log"
	"os"

	"github.com/darcioSoares/stark/internal/config"
	"github.com/darcioSoares/stark/internal/routes"
	"github.com/darcioSoares/stark/internal/services"
	
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)



func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	port := os.Getenv("PORT")
	config.LoadEnvVars()

	e := echo.New()

	//rotas
	routes.SetupRoutes(e)
	routes.SetupRoutesWebhook(e)
	
	//goroutine
	go func() {
		fmt.Println("Iniciando envio periódico de invoice...")
		services.SendRequestsEveryHour()
	}()

	log.Fatal(e.Start(":" + port))
}
