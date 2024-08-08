package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger() func(*fiber.Ctx) error {
	return logger.New(logger.Config{
		Format:     "${time} ${latency} ${status} ${method} ${path}?${queryParams} ${error}\n",
		TimeFormat: "2006-01-02 15:04:05",
		CustomTags: map[string]logger.LogFunc{
			"latency": func(output logger.Buffer, _ *fiber.Ctx, data *logger.Data, _ string) (int, error) {
				latency := data.Stop.Sub(data.Start)
				return output.WriteString(fmt.Sprintf("%5v", latency.Round(time.Microsecond)))
			},
		},
	})
}
