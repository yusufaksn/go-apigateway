package proxy

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

func GetService(c *fiber.Ctx, url string) error {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	req.Header.Set("X-Trace-ID", GenerateTraceId())

	err := fasthttp.Do(req, resp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating request")
	}

	c.Status(resp.StatusCode())
	return c.Send(resp.Body())
}

func PostService(c *fiber.Ctx, url string) error {

	requestBody := c.Body()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.SetBody(requestBody)
	req.Header.Set("X-Trace-ID", GenerateTraceId())

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed!",
		})
	}

	return c.Status(resp.StatusCode()).Send(resp.Body())
}

func PutService(c *fiber.Ctx, url string) error {

	requestBody := c.Body()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("PUT")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Trace-ID", GenerateTraceId())
	req.SetBody(requestBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed!",
		})
	}

	return c.Status(resp.StatusCode()).Send(resp.Body())
}

func DeleteService(c *fiber.Ctx, url string) error {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.Set("X-Trace-ID", GenerateTraceId())
	req.Header.SetMethod("DELETE")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed!",
		})
	}

	return c.Status(resp.StatusCode()).Send(resp.Body())
}

func GenerateTraceId() string {
	return uuid.New().String()
}
