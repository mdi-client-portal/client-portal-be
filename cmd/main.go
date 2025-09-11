package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/config"
)

func main() {
    cfg := config.Load()
    
	config.ConnectDB(cfg)
    
    app := fiber.New()
    app.Listen(":3000")
}