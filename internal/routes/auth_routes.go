package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/handler"
)

func RegisterAuthRoutes(router fiber.Router, h *handler.AuthHandler) {
	router.Post("/auth/signin", h.Signin)
	// router.Post("/register", h.Register)
}
