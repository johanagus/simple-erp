package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/service"
	"github.com/johanagus/simple-erp/pkg/response"
)

type SupplierHandler struct {
	service service.SupplierService
}

func NewSupplierHandler(service service.SupplierService) *SupplierHandler {
	return &SupplierHandler{service: service}
}

func (h *SupplierHandler) FindAll(c *fiber.Ctx) error {
	result, err := h.service.FindAll()
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "supplier tidak di temukan", err)
	}

	if result == nil {
		return response.Error(c, fiber.StatusNotFound, "supplier tidak di temukan", nil)
	}

	return response.Success(c, fiber.StatusOK, "berhasil mendapatkan data", result)

}

func (h *SupplierHandler) FindByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	result, err := h.service.FindByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "supplier id : "+strconv.Itoa(id)+" is not found", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "supplier id di temukan", result)
}

func (h *SupplierHandler) SaveSupplier(c *fiber.Ctx) error {
	supplier := &domain.Supplier{}
	if err := c.BodyParser(supplier); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", err)
	}

	_, err := h.service.SaveSupplier(supplier)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "failed to save supplier", err)
	}

	return response.Success(c, fiber.StatusCreated, "supplier berhasil disimpan", nil)
}

func (h *SupplierHandler) UpdateSupplier(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	supplier := &domain.Supplier{}
	if err := c.BodyParser(supplier); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", err)
	}

	h.service.UpdateSupplier(id, supplier)
	return response.Success(c, fiber.StatusOK, "supplier berhasil di update", nil)
}
