package proxy

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetService(c *fiber.Ctx, url string) error {

	resp, err := http.Get(url)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).SendString("Error creating request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).SendString("Error reading response body")
	}

	c.Status(resp.StatusCode)

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
