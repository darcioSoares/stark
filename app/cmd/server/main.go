package main

import (
	"fmt"
	"log"
	"os"

	"github.com/darcioSoares/stark/internal/routes"
	"github.com/darcioSoares/stark/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	port := os.Getenv("PORT")

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	routes.SetupRoutes(e)
	routes.SetupRoutesWebhook(e)

	// goroutine
	go func() {
		fmt.Println("Iniciando envio periódico...")
		services.SendRequestsEveryHour()
	}()

	log.Fatal(e.Start(":" + port))

}
