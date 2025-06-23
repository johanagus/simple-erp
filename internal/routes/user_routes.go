package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/handler"
	"github.com/johanagus/simple-erp/internal/middleware"
)

func RegisterUserRoutes(router fiber.Router, h *handler.UserHandler) {
	router.Get("/users", middleware.RequirePermission("get_user"), h.FindAll)
	router.Get("/user/:id", middleware.RequirePermission("get_user"), h.FindByID)
	router.Post("/user", middleware.RequirePermission("create_user"), h.SaveUser)
	router.Put("/user/:id", middleware.RequirePermission("update_user"), h.UpdateUser)
}
