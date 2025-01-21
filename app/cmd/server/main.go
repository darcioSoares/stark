package main

import (
	"log"
	"os"

	"github.com/darcioSoares/stark/internal/routes"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	port := os.Getenv("PORT")
		
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	routes.SetupRoutes(e)
	routes.SetupRoutesWebhook(e)

	log.Fatal(e.Start(":" + port))

}