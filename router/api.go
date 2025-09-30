package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router *fiber.App) {
	api := router.Group("/api")

	ClientRoutes(api)
	InvoiceRoutes(api)
}
