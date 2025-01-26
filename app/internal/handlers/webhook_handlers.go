package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/darcioSoares/stark/internal/config"
	"github.com/darcioSoares/stark/internal/models"
	"github.com/darcioSoares/stark/internal/services"

	"github.com/labstack/echo/v4"
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/event"
	"github.com/starkinfra/core-go/starkcore/user/project"
)

// Variável global para RabbitMQService
var RabbitMQ *services.RabbitMQService

type Message struct {
	Amount int    `json:"amount"`
	Name   string `json:"name"`
}

func HandleWebhook(exchangeName, routingKey, message string) {
	err := RabbitMQ.SendMessage(exchangeName, routingKey, message)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	} else {
		log.Printf("Message sent: %s", message)
	}
}

func WebhookHandler(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da requisição:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao ler o corpo da requisição"})
	}

	signature := c.Request().Header.Get("Digital-Signature")
	if signature == "" {
		fmt.Println("Cabeçalho Digital-Signature ausente")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cabeçalho Digital-Signature ausente"})
	}

	user := &project.Project{
		Id:          config.IDProject,
		PrivateKey:  config.PrivateKey,
		Environment: "sandbox",
	}

	starkbank.User = user
	event.Parse(string(body), signature, user)

	if err := json.Unmarshal(body, &models.RequestWebhook); err != nil {
		fmt.Println("Erro ao decodificar o evento:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao decodificar o evento"})
	}

	name := models.RequestWebhook.Event.Log.Invoice.Name
	amount := models.RequestWebhook.Event.Log.Invoice.Amount

	if models.RequestWebhook.Event.Subscription == "invoice" && models.RequestWebhook.Event.Log.Invoice.Status == "paid" {
		// fmt.Println("Invoice paga processada com sucesso!")

		// transfers, err := services.CreateTransfer(amount, name)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return c.JSON(http.StatusInternalServerError, map[string]string{
		// 		"error": err.Error(),
		// 	})
		// }

		// return c.JSON(http.StatusOK, transfers)

		message := Message{
			Amount: amount,
			Name:   name,
		}

		// serializar
		messageJSON, err := json.Marshal(message)
		if err != nil {
			log.Fatalf("Failed to serialize message to JSON: %v", err)
		}

		HandleWebhook("amq.fanout", "routing_key", string(messageJSON))

		return c.JSON(http.StatusOK, "Webhook enviado com sucesso")

	} else if models.RequestWebhook.Event.Subscription == "invoice" {
		fmt.Println("Invoice ainda não paga.")
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "Evento processado com sucesso"})
}
