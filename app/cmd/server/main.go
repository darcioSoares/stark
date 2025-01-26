package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	//"time"

	"github.com/darcioSoares/stark/internal/config"
	"github.com/darcioSoares/stark/internal/handlers"
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
	config.LoadEnvVars()

	e := echo.New()

	//rotas
	routes.SetupRoutes(e)
	routes.SetupRoutesWebhook(e)

	////////////////
	// Instanciar o serviço RabbitMQ
	rabbitmq := &services.RabbitMQService{}

	amqpURL := os.Getenv("AMQPURL")
	queueName := os.Getenv("QUEUENAME")
	exchangeName := os.Getenv("EXCHANGENAME")
	exchangeType := os.Getenv("EXCHANGETYPE")

	err = rabbitmq.Initialize(amqpURL, queueName, exchangeName, exchangeType)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	defer rabbitmq.Close()

	// Enviar mensagem para a exchange
	err = rabbitmq.SendMessage(exchangeName, "routing_key", "Hello, RabbitMQ!")
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	go func() {

		type MessageFilla struct {
			Amount int    `json:"amount"`
			Name   string `json:"name"`
		}

		messages, err := rabbitmq.ConsumeMessages()
		if err != nil {
			log.Fatalf("Failed to consume messages: %v", err)
		}

		for msg := range messages {

			var msgData MessageFilla
			err := json.Unmarshal(msg.Body, &msgData)
			if err != nil {
				log.Printf("Erro ao deserializar mensagem: %v", err)
				continue
			}

			transfers, err := services.CreateTransfer(msgData.Amount, msgData.Name)
			if err != nil {
				fmt.Println(err)
			}
			log.Printf("Mensagem recebida da fila e enviada para transfer : %v", transfers)
		}
	}()

	// Atribuir o RabbitMQService global no handler
	handlers.RabbitMQ = rabbitmq

	//goroutine
	go func() {
		fmt.Println("Iniciando envio periódico de invoice...")
		services.SendRequestsEveryHour()
	}()

	log.Fatal(e.Start(":" + port))
}
