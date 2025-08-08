package main

import (
	"GO_APIGATEWAY/db"
	"GO_APIGATEWAY/routes"
	"log"
	"net/http"
	"os"
	"time"

	fiberprometheus "github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	zap.L().Info("app starting...")

	/*err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}*/

	db.ConnectDB()

	tp, err := initTracer()
	if err != nil {
		log.Fatalf("tracer not started: %v", err)
	}
	defer func() { _ = tp.Shutdown(nil) }()

	http.DefaultTransport = otelhttp.NewTransport(http.DefaultTransport)
	http.DefaultClient = &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}

	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Concurrency:  256 * 1024,
	})

	app.Use(func(c *fiber.Ctx) error {
		tracer := otel.Tracer("api-gateway")
		ctx, span := tracer.Start(c.Context(), c.Method()+" "+c.Path())
		defer span.End()

		span.SetAttributes(
			attribute.String("http.method", c.Method()),
			attribute.String("http.url", c.OriginalURL()),
		)

		c.SetUserContext(ctx)
		return c.Next()
	})

	prometheus := fiberprometheus.New("fiber_app")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	routes.SetupRoutes(app)

	port := ":8080"
	log.Fatal(app.Listen(port))
}

func initTracer() (*sdktrace.TracerProvider, error) {
	jaegerURL := os.Getenv("JAEGER_ENDPOINT")
	if jaegerURL == "" {
		jaegerURL = "http://localhost:14268/api/traces"
	}

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerURL)))
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("api-gateway"),
			attribute.String("environment", "dev"),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}
