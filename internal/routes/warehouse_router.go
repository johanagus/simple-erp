package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/handler"
)

func RegisterWarehouseRoutes(router fiber.Router, h *handler.WarehouseHandler) {

	router.Get("/warehouses", h.FindAllWarehouses)
	router.Get("/warehouse/:id", h.FindWarehouseByID)
	router.Post("/warehouse", h.CreateWarehouse)
	router.Put("/warehouse/:id", h.UpdateWarehouse)
}
