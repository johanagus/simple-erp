package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/johanagus/simple-erp/config"
	"github.com/johanagus/simple-erp/internal/handler"
	"github.com/johanagus/simple-erp/internal/middleware"
	"github.com/johanagus/simple-erp/internal/repository"
	"github.com/johanagus/simple-erp/internal/routes"
	"github.com/johanagus/simple-erp/internal/service"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	DB := config.InitDB()

	config.SeedData(DB) // Seed initial data

	// Init handler, service, repo...
	authRepo := repository.NewAuthRepository(DB)
	roleRepo := repository.NewRoleRepository(DB)
	authService := service.NewAuthService(authRepo, roleRepo)
	authHandler := handler.NewAuthHandler(authService)

	userRepo := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	categoryRepo := repository.NewCategoryRepository(DB)
	productRepo := repository.NewProductRepository(DB)
	productService := service.NewProductService(productRepo, categoryRepo)
	productHandler := handler.NewProductHandler(productService)

	warehouseRepo := repository.NewWarehouseRepository(DB)
	warehouseService := service.NewWarehouseService(warehouseRepo)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)

	supplierRepo := repository.NewSupplierRepository(DB)
	supplierService := service.NewSupplierService(supplierRepo)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	storeRepo := repository.NewStoreRepository(DB)
	storeService := service.NewStoreService(storeRepo)
	storeHandler := handler.NewStoreHandler(storeService)

	customerRepo := repository.NewCustomerRepository(DB)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	app.Use(middleware.Logger())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // frontend origin
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Register routes
	routes.RegisterRoutes(app, routes.RouteConfig{
		AuthHandler:      authHandler,
		UserHandler:      userHandler,
		ProductHandler:   productHandler,
		WarehouseHandler: warehouseHandler,
		SupplierHandler:  supplierHandler,
		StoreHandler:     storeHandler,
		CustomerHandler:  customerHandler,
	})

	app.Listen(":8000")
}
