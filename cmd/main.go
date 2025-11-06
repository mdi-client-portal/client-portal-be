package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/internal/jobs"

	"github.com/robfig/cron/v3"
	// "github.com/mdi-client-portal/client-portal-be/database/seeders"
	"github.com/mdi-client-portal/client-portal-be/router"
)

func main() {
	config.LoggerInit()
	config.EnvInit()

	config.Log.Info("Starting application...")

	defer config.Log.Sync()

	config.ConnectDB()

	// seeders.ClientSeeder(config.DB)
	// seeders.InvoiceSeeder(config.DB)

	config.Log.Info("Setting up cron jobs...")
	cronJob := cron.New()
	cronJob.AddFunc("* * * * *", func() {jobs.EmailCron(config.DB)})
	cronJob.Start()
	config.Log.Info("Cron jobs set up successfully")

	config.Log.Info("Setting up routes and starting server...")
	app := fiber.New()
	app.Use(cors.New())
	router.SetupRoutes(app)

	config.Log.Info("Server is running on port " + config.Env.Port)
	app.Listen("0.0.0.0:" + config.Env.Port)
}
