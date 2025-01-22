package handlers

import (
	"fmt"
	"net/http"

	_ "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type RequestTransactionBody struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func Welcome1(c echo.Context) error {
	return c.JSON(http.StatusOK, "Welcome api")
}

func GetReturn1(c echo.Context) error {

	body := new(RequestBody)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Erro ao processar o corpo da requisição",
		})
	}

	if err := c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":  "Dados inválidos",
			"detail": err.Error(),
		})
	}

	fmt.Println("Corpo da requisição:", body)

	return c.JSON(http.StatusOK, body)
}
