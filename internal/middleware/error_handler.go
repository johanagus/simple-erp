package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/pkg/response"
)

func ErrorHandler(c *fiber.Ctx, err error) error {

	// Jika error sudah bertipe *fiber.Error
	if e, ok := err.(*fiber.Error); ok {
		return response.Error(c, e.Code, e.Message, nil)
	}

	// Log error untuk server error
	log.Printf("Server Error: %v\n", err)

	// Default internal server error
	return response.Error(c, fiber.StatusInternalServerError, "Terjadi kesalahan pada server", nil)
}
