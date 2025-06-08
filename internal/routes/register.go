package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/johanagus/simple-erp/internal/handler"
	"github.com/johanagus/simple-erp/internal/middleware"
)

type RouteConfig struct {
	AuthHandler *handler.AuthHandler
	UserHandler *handler.UserHandler
	// Tambah handler lain di sini
}

func RegisterRoutes(app *fiber.App, cfg RouteConfig) {
	api := app.Group("/api/v1")

	//public route
	RegisterAuthRoutes(api, cfg.AuthHandler)

	// protected route
	protected := api.Group("")
	protected.Use(middleware.JWTProtected())

	RegisterUserRoutes(protected, cfg.UserHandler)

}
