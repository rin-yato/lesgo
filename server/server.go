package server

import (
	"github.com/gofiber/fiber/v2"

	"barcode-api/database"
	"barcode-api/server/middleware"
)

type FiberServer struct {
	*fiber.App
	DB database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			AppName: "barcode-api",
		}),

		DB: database.New(),
	}

	middleware.RegisterMiddleware(server.App)

	return server
}
