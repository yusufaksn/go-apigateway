package auth

import (
	"context"
	"log"
	"os"
	"time"

	"GO_APIGATEWAY/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5" // v5 sürümünü kullanıyoruz
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("APP_KEY")) // JWT için secret key

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterUser(c *fiber.Ctx) error {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz giriş"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Hash error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Hash error"})
	}

	_, err = db.DB.Exec(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, string(hashedPassword))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı veritabanına eklenemedi"})
	}

	return c.JSON(fiber.Map{"message": "Successfull!"})
}

func LoginUser(c *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Fail"})
	}

	var storedPassword string
	err := db.DB.QueryRow(context.Background(), "SELECT password FROM users WHERE username=$1", loginData.Username).Scan(&storedPassword)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Fail"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Fail"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(loginData.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Fail"})
	}

	claims := &jwt.RegisteredClaims{
		Issuer:    loginData.Username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Fail"})
	}

	return c.JSON(fiber.Map{
		"message": "Successfull",
		"token":   signedToken,
	})
}
