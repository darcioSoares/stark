package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/darcioSoares/stark/internal/services"
    "github.com/darcioSoares/stark/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/event"
	"github.com/starkinfra/core-go/starkcore/user/project"
)

var privateKeyContent = `-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIN0NFH1lGEzLXhnaXxKKBqC3J1WWuLtiRAzSEfRXBqTgoAcGBSuBBAAK
oUQDQgAEu4gONKh9t794DaLahDib/rfL5aGyR0V/0RSvZ6cd46y/j78ybFWsd04Y
kiDAFLMFGeLuP0u4n2JV1JIPyBSL6w==
-----END EC PRIVATE KEY-----`

var user = &project.Project{
	Id:          "6250122287513600",
	PrivateKey:  privateKeyContent,
	Environment: "sandbox",
}

func WebhookHandler(c echo.Context) error {
	// Lê o corpo da requisição
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da requisição:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao ler o corpo da requisição"})
	}

	// Obtenha o cabeçalho de assinatura
	signature := c.Request().Header.Get("Digital-Signature")
	if signature == "" {
		fmt.Println("Cabeçalho Digital-Signature ausente")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cabeçalho Digital-Signature ausente"})
	}

	starkbank.User = user

	event.Parse(string(body), signature, user)

	// var RequestWebhook struct {
	// 	Event struct {
	// 		Created      string `json:"created"`
	// 		ID           string `json:"id"`
	// 		Subscription string `json:"subscription"`
	// 		WorkspaceID  string `json:"workspaceId"`
	// 		Log          struct {
	// 			Type    string `json:"type"`
	// 			Created string `json:"created"`
	// 			Invoice struct {
	// 				ID         string `json:"id"`
	// 				Status     string `json:"status"`
	// 				Amount     int    `json:"amount"`
	// 				Name       string `json:"name"`
	// 				TaxID      string `json:"taxId"`
	// 				Created    string `json:"created"`
	// 				Nominal    int    `json:"nominalAmount"`
	// 				Link       string `json:"link"`
	// 				PDF        string `json:"pdf"`
	// 				Expiration int    `json:"expiration"`
	// 			} `json:"invoice"`
	// 		} `json:"log"`
	// 	} `json:"event"`
	// }

	// Decodifica o evento no formato esperado
	if err := json.Unmarshal(body, &models.RequestWebhook); err != nil {
		fmt.Println("Erro ao decodificar o evento:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao decodificar o evento"})
	}

	// Loga os detalhes do evento no console
	fmt.Println("Evento Recebido:")
	fmt.Printf("ID do Evento: %s\n", models.RequestWebhook.Event.ID)
	fmt.Printf("Tipo de Assinatura: %s\n", models.RequestWebhook.Event.Subscription)
	fmt.Printf("Tipo de Log: %s\n", models.RequestWebhook.Event.Log.Type)
	fmt.Printf("Status da Invoice: %s\n", models.RequestWebhook.Event.Log.Invoice.Status)
	fmt.Printf("Nome: %s, Valor: %d, PDF: %s\n", models.RequestWebhook.Event.Log.Invoice.Name, models.RequestWebhook.Event.Log.Invoice.Amount, models.RequestWebhook.Event.Log.Invoice.PDF)

	name := models.RequestWebhook.Event.Log.Invoice.Name
	amount := models.RequestWebhook.Event.Log.Invoice.Amount

	if models.RequestWebhook.Event.Subscription == "invoice" && models.RequestWebhook.Event.Log.Invoice.Status == "paid" {
		fmt.Println("Invoice paga processada com sucesso!")

		transfers, err := services.CreateTransfer(amount, name)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		// Retorna a resposta para o cliente
		return c.JSON(http.StatusOK, transfers)

	} else if models.RequestWebhook.Event.Subscription == "invoice" {
		fmt.Println("Invoice ainda não paga ou outro tipo de evento.")
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "Evento processado com sucesso"})
}
