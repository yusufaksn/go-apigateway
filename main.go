package main

import (
	"GO_APIGATEWAY/db"
	"GO_APIGATEWAY/routes"
	"log"
	"time"

	fiberprometheus "github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {

	defer zap.L().Sync()
	zap.L().Info("app starting...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()

	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Concurrency:  256 * 1024,
	})

	prometheus := fiberprometheus.New("fiber_app")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	routes.SetupRoutes(app)

	port := ":8080"
	log.Fatal(app.Listen(port))

}
