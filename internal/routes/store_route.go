package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/handler"
	"github.com/johanagus/simple-erp/internal/middleware"
)

func RegisterStoreRoutes(router fiber.Router, h *handler.StoreHandler) {
	router.Get("/stores", middleware.RequirePermission("get_store"), h.FindAll)
	router.Get("/store/:id", middleware.RequirePermission("get_store"), h.FindByID)
	router.Post("/store", middleware.RequirePermission("create_store"), h.CreateStore)
	router.Put("/store/:id", middleware.RequirePermission("update_store"), h.UpdateStore)
}
