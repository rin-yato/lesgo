package health

import (
	"barcode-api/database"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	db database.Service
}

func NewHealthHandler(db database.Service) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

func (h *HealthHandler) Ping(c *fiber.Ctx) error {
	return c.SendString("Pong")
}

func (h *HealthHandler) DBHealth(c *fiber.Ctx) error {
	return c.JSON(h.db.Health())
}
