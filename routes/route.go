package routes

import (
	"GO_APIGATEWAY/handlers/a"
	"GO_APIGATEWAY/handlers/auth"
	"GO_APIGATEWAY/handlers/b"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your_secret_key")

func SetupRoutes(app *fiber.App) {
	// Service A Routes
	app.Get("/service-a/*", authMiddleware, a.GetMethod)
	app.Post("/service-a/*", a.PostMethod)
	app.Put("/service-a/*", a.PutMethod)
	app.Delete("/service-a/*", a.DeleteMethod)

	// Service B Routes
	app.Get("/service-b/*", b.GetMethod)
	app.Post("/service-b/*", b.PostMethod)
	app.Put("/service-b/*", b.PutMethod)
	app.Delete("/service-b/*", b.DeleteMethod)

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/register", auth.RegisterUser)
	app.Post("/login", auth.LoginUser)

	app.Get("/protected", authMiddleware, protectedRoute)
}

func authMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	tokenString = tokenString[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	return c.Next()
}

func protectedRoute(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Welcome to the protected route!"})
}
