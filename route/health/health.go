package health

import "barcode-api/server"

func RegisterRoutes(server *server.FiberServer) {
	handler := NewHealthHandler(server.DB)

	group := server.Group("/health")

	group.Get("/ping", handler.Ping)
	group.Get("/db", handler.DBHealth)
}
