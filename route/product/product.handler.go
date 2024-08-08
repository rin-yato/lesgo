package product

import (
	"barcode-api/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	db database.Service
}

func NewProductHandler(db database.Service) *ProductHandler {
	return &ProductHandler{db: db}
}

func (u *ProductHandler) GetAll(c *fiber.Ctx) error {
	return c.SendString("Get All Products")
}

func (u *ProductHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"id": id, "type": "Get Product By ID"})
}

func (u *ProductHandler) Create(c *fiber.Ctx) error {
	payload := new(struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	})

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"name": payload.Name, "price": payload.Price})
}

func (u *ProductHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	payload := new(struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	})

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"id": id, "name": payload.Name, "price": payload.Price})
}

func (u *ProductHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"id": id, "type": "Delete Product"})
}
