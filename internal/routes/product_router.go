package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/handler"
)

func RegisterProductRoutes(router fiber.Router, h *handler.ProductHandler) {
	router.Get("/products", h.FindAll)
	router.Get("/product/:id", h.FindByID)
	router.Get("/product/sku/:sku", h.FindBySKU)
	router.Get("/product/barcode/:barcode", h.FindByBarcode)
	router.Get("/products/search", h.Search)
	router.Post("/product", h.SaveProduct)
	router.Put("/product", h.UpdateProduct)
}
