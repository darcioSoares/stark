package routes

import (
	"github.com/darcioSoares/stark/internal/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutesWebhook(e *echo.Echo) {
	api := e.Group("/webhook")

	api.POST("", handlers.WebhookHandler)

}
