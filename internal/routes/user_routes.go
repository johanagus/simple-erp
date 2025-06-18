package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/handler"
	"github.com/johanagus/simple-erp/internal/middleware"
)

func RegisterUserRoutes(router fiber.Router, h *handler.UserHandler) {
	router.Get("/users", middleware.RequirePermission("get_users"), h.FindAll)
	router.Get("/user/:id", middleware.RequirePermission("get_user_by_id"), h.FindByID)
	router.Post("/user", middleware.RequirePermission("create_user"), h.SaveUser)
}
