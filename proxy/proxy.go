package proxy

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetService(c *fiber.Ctx, url string) error {
	// HTTP GET isteği yap ve yanıtı al
	resp, err := http.Get(url)
	if err != nil {
		// Hata durumunda iç hatayı gönder

		return c.Status(fiber.StatusInternalServerError).SendString("Error creating request")
	}
	defer resp.Body.Close()

	// Gelen yanıtı okuma
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Yanıtı okuma hatası

		return c.Status(fiber.StatusInternalServerError).SendString("Error reading response body")
	}

	// Yanıtın HTTP durum kodunu istemciye ilet
	c.Status(resp.StatusCode)

	// Gelen yanıtı Fiber üzerinden ilet
	return c.Send(body)
}

func PostService(c *fiber.Ctx, url string) error {
	return c.JSON(fiber.Map{"method post proxy": true})
}

func PutService(c *fiber.Ctx, url string) error {
	return c.JSON(fiber.Map{"method put proxy": true})
}

func DeleteService(c *fiber.Ctx, url string) error {
	return c.JSON(fiber.Map{"method put proxy": true})
}
