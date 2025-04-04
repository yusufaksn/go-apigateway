package b

import (
	"GO_APIGATEWAY/proxy"

	"github.com/gofiber/fiber/v2"
)

var url = "http://localhost:8181"

func GetMethod(c *fiber.Ctx) error {
	return proxy.GetService(c, url)
}

func PostMethod(c *fiber.Ctx) error {
	return proxy.PostService(c, url)
}

func PutMethod(c *fiber.Ctx) error {
	return proxy.PutService(c, url)
}

func DeleteMethod(c *fiber.Ctx) error {
	return proxy.DeleteService(c, url)
}
