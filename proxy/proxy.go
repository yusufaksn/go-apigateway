package proxy

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sony/gobreaker"
	"github.com/valyala/fasthttp"
)

var cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
	Name:        "Service Circuit Breaker",
	Timeout:     10 * 1000000000, // 30 seconds
	MaxRequests: 3,               // Max requests in the half-open state before trying
})

func GetService(c *fiber.Ctx, url string) error {
	result, err := cb.Execute(func() (interface{}, error) {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		// Get the path after /service-a or /service-b
		path := c.Path()
		parts := strings.Split(path, "/")
		if len(parts) > 2 {
			path = "/" + strings.Join(parts[2:], "/")
		}

		targetURL := url + path
		fmt.Printf("Target URL: %s\n", targetURL)

		req.SetRequestURI(targetURL)
		req.Header.SetMethod("GET")
		req.Header.Set("X-Trace-ID", GenerateTraceId())

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		err := fasthttp.Do(req, resp)
		if err != nil {
			return nil, err
		}
		copiedResp := &fasthttp.Response{}
		resp.CopyTo(copiedResp)
		fasthttp.ReleaseResponse(resp)
		fmt.Printf("Response Status: %d, Body: %s\n", resp.StatusCode(), string(resp.Body()))
		return copiedResp, nil
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating request: " + err.Error())
	}
	res := result.(*fasthttp.Response)
	return c.Status(res.StatusCode()).Send(res.Body())
}

func PostService(c *fiber.Ctx, url string) error {
	result, err := cb.Execute(func() (interface{}, error) {
		requestBody := c.Body()
		fmt.Printf("Request Body: %s\n", string(requestBody))

		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		path := c.Path()
		parts := strings.Split(path, "/")
		if len(parts) > 2 {
			path = "/" + strings.Join(parts[2:], "/")
		}

		targetURL := url + path
		fmt.Printf("Target URL: %s\n", targetURL)

		req.SetRequestURI(targetURL)
		req.Header.SetMethod("POST")
		req.Header.Set("Content-Type", "application/json")
		req.SetBody(requestBody)
		req.Header.Set("X-Trace-ID", GenerateTraceId())

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		err := fasthttp.Do(req, resp)
		if err != nil {
			return nil, err
		}
		copiedResp := &fasthttp.Response{}
		resp.CopyTo(copiedResp)
		fasthttp.ReleaseResponse(resp)
		fmt.Printf("Response Status: %d, Body: %s\n", resp.StatusCode(), string(resp.Body()))
		return copiedResp, nil
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating request: " + err.Error())
	}
	res := result.(*fasthttp.Response)
	return c.Status(res.StatusCode()).Send(res.Body())
}

func PutService(c *fiber.Ctx, url string) error {
	result, err := cb.Execute(func() (interface{}, error) {
		requestBody := c.Body()
		fmt.Printf("Request Body: %s\n", string(requestBody))

		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		path := c.Path()
		parts := strings.Split(path, "/")
		if len(parts) > 2 {
			path = "/" + strings.Join(parts[2:], "/")
		}

		targetURL := url + path
		fmt.Printf("Target URL: %s\n", targetURL)

		req.SetRequestURI(targetURL)
		req.Header.SetMethod("PUT")
		req.Header.Set("Content-Type", "application/json")
		req.SetBody(requestBody)
		req.Header.Set("X-Trace-ID", GenerateTraceId())

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		err := fasthttp.Do(req, resp)
		if err != nil {
			return nil, err
		}
		copiedResp := &fasthttp.Response{}
		resp.CopyTo(copiedResp)
		fasthttp.ReleaseResponse(resp)
		fmt.Printf("Response Status: %d, Body: %s\n", resp.StatusCode(), string(resp.Body()))
		return copiedResp, nil
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating request: " + err.Error())
	}
	res := result.(*fasthttp.Response)
	return c.Status(res.StatusCode()).Send(res.Body())
}

func DeleteService(c *fiber.Ctx, url string) error {
	result, err := cb.Execute(func() (interface{}, error) {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		path := c.Path()
		parts := strings.Split(path, "/")
		if len(parts) > 2 {
			path = "/" + strings.Join(parts[2:], "/")
		}

		targetURL := url + path
		fmt.Printf("Target URL: %s\n", targetURL)

		req.SetRequestURI(targetURL)
		req.Header.SetMethod("DELETE")
		req.Header.Set("X-Trace-ID", GenerateTraceId())

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		err := fasthttp.Do(req, resp)
		if err != nil {
			return nil, err
		}
		copiedResp := &fasthttp.Response{}
		resp.CopyTo(copiedResp)
		fasthttp.ReleaseResponse(resp)
		fmt.Printf("Response Status: %d, Body: %s\n", resp.StatusCode(), string(resp.Body()))
		return copiedResp, nil
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating request: " + err.Error())
	}
	res := result.(*fasthttp.Response)
	return c.Status(res.StatusCode()).Send(res.Body())
}

func GenerateTraceId() string {
	return uuid.New().String()
}
