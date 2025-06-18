package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/handler"
	"github.com/johanagus/simple-erp/internal/middleware"
)

func RegisterProductRoutes(router fiber.Router, h *handler.ProductHandler) {
	router.Get("/products", middleware.RequirePermission("get_products"), h.FindAll)
	router.Get("/product/:id", middleware.RequirePermission("get_product_by_id"), h.FindByID)
	router.Get("/product/sku/:sku", middleware.RequirePermission("get_product_by_sku"), h.FindBySKU)
	router.Get("/product/barcode/:barcode", middleware.RequirePermission("get_product_by_barcode"), h.FindByBarcode)
	router.Get("/products/search", middleware.RequirePermission("search_product"), h.Search)
	router.Post("/product", middleware.RequirePermission("create_product"), h.SaveProduct)
	router.Put("/product", middleware.RequirePermission("update_products"), h.UpdateProduct)
}
