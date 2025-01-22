package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/darcioSoares/stark/internal/models"
)



func GetReturnPay(c echo.Context) error {
	
	var payload models.RequestWebhooktPayload

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

	return c.JSON(http.StatusOK, nil)
}
