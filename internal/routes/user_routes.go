package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/handler"
)

func RegisterUserRoutes(router fiber.Router, h *handler.UserHandler) {
	router.Get("/users", h.FindAll)
	router.Get("/user/:id", h.FindByID)
	router.Post("/user", h.SaveUser)
}
