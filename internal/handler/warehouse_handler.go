package handler

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/service"
	"github.com/johanagus/simple-erp/pkg/response"
)

type WarehouseHandler struct {
	Service service.WarehouseService
}

func NewWarehouseHandler(service service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{Service: service}
}

func (h *WarehouseHandler) FindAllWarehouses(c *fiber.Ctx) error {
	warehouses, err := h.Service.FindAllWarehouses()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "failed to retrieve warehouses", err)
	}

	if len(warehouses) == 0 {
		return response.Error(c, fiber.StatusNotFound, "no warehouses found", nil)
	}

	return response.Success(c, fiber.StatusOK, "warehouses retrieved successfully", warehouses)
}

func (h *WarehouseHandler) FindWarehouseByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid warehouse ID", err)
	}

	warehouse, err := h.Service.FindWarehouseByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "warehouse not found", err)
	}

	return response.Success(c, fiber.StatusOK, "warehouse retrieved successfully", warehouse)
}

func (h *WarehouseHandler) CreateWarehouse(c *fiber.Ctx) error {
	warehouse := new(domain.Warehouse)
	if err := c.BodyParser(warehouse); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(warehouse); err != nil {
		var validationErrors []string
		for _, fieldErr := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fieldErr.Field()+": "+fieldErr.ActualTag())
		}
		return response.Error(c, fiber.StatusBadRequest, "validation failed", validationErrors)
	}

	id, err := h.Service.CreateWarehouse(warehouse)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "failed to create warehouse", err)
	}

	return response.Success(c, fiber.StatusCreated, "warehouse created successfully", fiber.Map{"id": id})
}

func (h *WarehouseHandler) UpdateWarehouse(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid warehouse ID", err)
	}

	warehouse := new(domain.Warehouse)
	if err := c.BodyParser(warehouse); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", err)
	}
	warehouse.ID = uint(id)

	updatedWarehouse, err := h.Service.UpdateWarehouse(warehouse)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "failed to update warehouse", err)
	}

	return response.Success(c, fiber.StatusOK, "warehouse updated successfully", updatedWarehouse)
}
