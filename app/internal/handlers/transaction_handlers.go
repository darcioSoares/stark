package handlers

import (
	"fmt"
	"net/http"

	"github.com/darcioSoares/stark/internal/services"
	"github.com/labstack/echo/v4"
)

func StoreInvoiceHandler(c echo.Context) error {
	invoices, err := services.CreateInvoice()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// Retorna a resposta para o cliente
	return c.JSON(http.StatusOK, invoices)
}

func CreateTransferHandler(c echo.Context) error {
	transfers, err := services.CreateTransfer()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// Retorna a resposta para o cliente
	return c.JSON(http.StatusOK, transfers)
}