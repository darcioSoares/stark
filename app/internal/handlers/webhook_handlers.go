package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Event struct {
	ID           string `json:"id"`
	IsDelivered  bool   `json:"isDelivered"`
	Subscription string `json:"subscription"`
	Created      string `json:"created"`
	Log          struct {
		ID       string   `json:"id"`
		Errors   []string `json:"errors"`
		Type     string   `json:"type"`
		Created  string   `json:"created"`
		Transfer struct {
			ID             string   `json:"id"`
			Status         string   `json:"status"`
			Amount         int      `json:"amount"`
			Name           string   `json:"name"`
			BankCode       string   `json:"bankCode"`
			BranchCode     string   `json:"branchCode"`
			AccountNumber  string   `json:"accountNumber"`
			TaxID          string   `json:"taxId"`
			Tags           []string `json:"tags"`
			Created        string   `json:"created"`
			Updated        string   `json:"updated"`
			TransactionIds []string `json:"transactionIds"`
			Fee            int      `json:"fee"`
		} `json:"transfer"`
	} `json:"log"`
}

type RequestPayload struct {
	Event Event `json:"event"`
}

func GetReturnPay(c echo.Context) error {
	
	var payload RequestPayload

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	fmt.Printf("Corpo da requisição: %+v\n", payload)

	// Retorna uma resposta de sucesso
	// return c.JSON(http.StatusOK, map[string]string{
	// 	"message": "Request received successfully",
	// })

	return c.JSON(http.StatusOK, payload)
}
