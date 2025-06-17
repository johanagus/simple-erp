package handler

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/service"
	"github.com/johanagus/simple-erp/pkg/response"
)

type InventoriHandler struct {
	service service.InventoriService
}

func NewInventoriHandler(service service.InventoriService) *InventoriHandler {
	return &InventoriHandler{
		service: service,
	}
}

// Define methods for the InventoriHandler to handle requests
func (h *InventoriHandler) FindInventoriByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID parameter")
	}
	// Call the service to find the inventory by ID
	inventori, err := h.service.FindByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Inventory not found")
	}

	return response.Success(c, fiber.StatusOK, "Inventory found", inventori)
}

func (h *InventoriHandler) FindInventoriByWarehouseID(c *fiber.Ctx) error {
	whID, err := c.ParamsInt("warehouse_id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Warehouse ID parameter")
	}

	inventories, err := h.service.FindByWHID(whID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "No inventory found for this warehouse")
	}

	return response.Success(c, fiber.StatusOK, "Inventories found", inventories)
}

func (h *InventoriHandler) SaveInventori(c *fiber.Ctx) error {
	var inventori *domain.Inventory
	if err := c.BodyParser(&inventori); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	validate := validator.New()
	if err := validate.Struct(inventori); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid inventory data", err.Error())
	}

	success, err := h.service.SaveInventori(inventori)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to save inventory", err.Error())
	}

	if !success {
		return response.Error(c, fiber.StatusConflict, "Failed to save inventory", nil)
	}

	return response.Success(c, fiber.StatusCreated, "Inventory saved successfully", inventori)
}

func (h *InventoriHandler) UpdateInventori(c *fiber.Ctx) error {
	var inventori *domain.Inventory
	if err := c.BodyParser(&inventori); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	updatedInventori, err := h.service.UpdateInventori(inventori)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update inventory", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Inventory updated successfully", updatedInventori)
}
