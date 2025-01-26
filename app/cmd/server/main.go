package main

import (
	"fmt"
	"log"
	"os"

	//"time"

	"github.com/darcioSoares/stark/internal/config"
	"github.com/darcioSoares/stark/internal/handlers"
	"github.com/darcioSoares/stark/internal/routes"
	"github.com/darcioSoares/stark/internal/services"

	//"github.com/streadway/amqp"

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

	e.Validator = &CustomValidator{validator: validator.New()}

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

	// // Consumir mensagens
	// messages, err := rabbitmq.ConsumeMessages()
	// if err != nil {
	// 	log.Fatalf("Failed to consume messages: %v", err)
	// }

	// // Processar mensagens recebidas
	// for msg := range messages {
	// 	log.Printf("Received message: %s", string(msg.Body))
	// }

	// Atribuir o RabbitMQService global no handler
	handlers.RabbitMQ = rabbitmq

	// Usar o handler para enviar uma mensagem
	handlers.HandleWebhook(exchangeName, "routing_key", "Hello, RabbitMQ  hellllooo!")

	//////////////////////////

	//goroutine
	go func() {
		fmt.Println("Iniciando envio periódico de invoice...")
		services.SendRequestsEveryHour()
	}()

	log.Fatal(e.Start(":" + port))
}
