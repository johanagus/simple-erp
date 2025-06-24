package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/handler"
	"github.com/johanagus/simple-erp/internal/middleware"
)

func RegisterCustomerRoutes(router fiber.Router, h *handler.CustomerHandler) {
	router.Get("/customers", middleware.RequirePermission("get_customer"), h.FindAll)
	router.Get("/customer/:id", middleware.RequirePermission("get_customer"), h.FindByID)
	router.Post("/customer", middleware.RequirePermission("create_customer"), h.CreateCustomer)
	router.Put("/customer/:id", middleware.RequirePermission("update_customer"), h.UpdateCustomer)
}
