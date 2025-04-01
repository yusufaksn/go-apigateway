package main

import (
	"GO_APIGATEWAY/routes"
	"log"
	"time"

	"GO_APIGATEWAY/db"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Kullanıcı modeli
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// PostgreSQL bağlantısını açan fonksiyon

func main() {

	db.ConnectDB()

	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Concurrency:  256 * 1024,
	})

	defer zap.L().Sync()
	zap.L().Info("app starting...")

	routes.SetupRoutes(app)

	port := ":8080"
	log.Fatal(app.Listen(port))

}
