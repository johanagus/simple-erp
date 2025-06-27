package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/handler"
	"github.com/johanagus/simple-erp/internal/middleware"
)

func RegisterWarehouseRoutes(router fiber.Router, h *handler.WarehouseHandler) {

	router.Get("/warehouses", middleware.RequirePermission("get_warehouse"), h.FindAllWarehouses)
	router.Get("/warehouse/:id", middleware.RequirePermission("get_warehouse"), h.FindWarehouseByID)
	router.Post("/warehouse", middleware.RequirePermission("create_warehouse"), h.CreateWarehouse)
	router.Put("/warehouse/:id", middleware.RequirePermission("update_warehouse"), h.UpdateWarehouse)
}
