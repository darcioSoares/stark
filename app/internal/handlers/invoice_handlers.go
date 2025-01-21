package handlers

import (
	"fmt"
	"net/http"
	
	"github.com/labstack/echo/v4"
	_"github.com/go-playground/validator/v10"
)

type RequestBody struct {
	Name  string `json:"name" validate:"required"`    
	Email string `json:"email" validate:"required,email"` 
}
func GetReturn(c echo.Context) error {

	body := new(RequestBody)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Erro ao processar o corpo da requisição",
		})
	}

	if err := c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Dados inválidos",
			"detail": err.Error(),
		})
	}

	fmt.Println("Corpo da requisição:", body)

	return c.JSON(http.StatusOK, body)
}


