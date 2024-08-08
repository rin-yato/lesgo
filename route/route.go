package route

import (
	"barcode-api/route/health"
	"barcode-api/route/product"
	"barcode-api/server"
)

func RegisterFiberRoutes(server *server.FiberServer) {
	health.RegisterRoutes(server)
	product.RegisterRoutes(server)
}
