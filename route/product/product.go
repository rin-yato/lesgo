package product

import (
	"barcode-api/server"
)

func RegisterRoutes(server *server.FiberServer) {
	handler := NewProductHandler(server.DB)

	group := server.Group("/product")

	group.Get("/", handler.GetAll)
	group.Get("/:id", handler.GetByID)

	group.Post("/", handler.Create)
	group.Put("/:id", handler.Update)
	group.Delete("/:id", handler.Delete)
}
