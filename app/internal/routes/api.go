package routes

import (
	"github.com/darcioSoares/stark/internal/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	api := e.Group("/api")

	api.GET("/", handlers.Welcome)
	api.POST("/invoices", handlers.GetReturn)

}
