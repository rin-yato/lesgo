package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RegisterMiddleware(server *fiber.App) {
	// One-liner
	server.Use(recover.New())

	// Custom Configuration
	server.Use(Logger())
}
