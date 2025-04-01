package b

import (
	"github.com/gofiber/fiber/v2"
)

func GetMethod(c *fiber.Ctx) error {
	// Construct the URL for service A
	/*url := "https://api.spacexdata.com/v4/launches" + c.Params("*")
	return proxy.GetService(c, url)*/

	return c.JSON(fiber.Map{"method": "post"})

}

func PostMethod(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"method": "post"})
}

func PutMethod(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"method": "put"})
}

func DeleteMethod(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"method": "delete"})
}
