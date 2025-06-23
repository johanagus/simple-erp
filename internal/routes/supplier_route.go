package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/handler"
	"github.com/johanagus/simple-erp/internal/middleware"
)

func RegisterSupplierRoutes(router fiber.Router, h *handler.SupplierHandler) {
	router.Get("/suppliers", middleware.RequirePermission("get_supplier"), h.FindAll)
	router.Get("/supplier/:id", middleware.RequirePermission("get_supplier"), h.FindByID)
	router.Post("/supplier", middleware.RequirePermission("create_supplier"), h.SaveSupplier)
	router.Put("/supplier/:id", middleware.RequirePermission("update_supplier"), h.UpdateSupplier)
}
