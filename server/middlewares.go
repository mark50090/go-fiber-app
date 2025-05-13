package servers

import (
	"fmt"
	"go-fiber-app/configs"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// middlewareLogger is a middleware that logs the request and response
func (s *server) middlewareLogger(ctx *fiber.Ctx) error {

	startTime := time.Now()

	err := ctx.Next()

	duration := time.Since(startTime)
	s.log.Info(fmt.Sprintf("[Out] %s %d %s (Duration: %s)",
		ctx.Method(),
		ctx.Response().StatusCode(),
		ctx.Path(),
		duration.String(),
	))
	return err
}

func (s *server) Cors(cfg *configs.Config) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Set allowed origins
		corsConfig := cors.Config{
			AllowOrigins: cfg.App.Cors.AllowOrigins,
			AllowHeaders: "Content-Type, Authorization",
			AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
		}

		// Apply CORS headers
		return cors.New(corsConfig)(c)
	}
}

func (s *server) statusColor(status int) string {
	if status >= 200 && status < 300 {
		return configs.Green
	}
	return configs.Reset
}
