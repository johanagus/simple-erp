package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/service"
	"github.com/johanagus/simple-erp/pkg/response"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) FindAll(c *fiber.Ctx) error {
	products, err := h.service.FindAll()
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "product tidak di temukan", err)
	}

	if len(products) == 0 {
		return response.Error(c, fiber.StatusNotFound, "product tidak di temukan", nil)
	}

	return response.Success(c, fiber.StatusOK, "berhasil mendapatkan data", products)
}

func (h *ProductHandler) FindByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product, err := h.service.FindByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "product id : "+strconv.Itoa(id)+" is not found", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "product id di temukan", product)
}

func (h *ProductHandler) FindBySKU(c *fiber.Ctx) error {
	product, err := h.service.FindBySKU(c.Params("sku"))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "product sku : "+c.Params("sku")+" is not found", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "product sku di temukan", product)
}

func (h *ProductHandler) FindByBarcode(c *fiber.Ctx) error {
	product, err := h.service.FindByBarcode(c.Params("barcode"))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "product barcode : "+c.Params("barcode")+" is not found", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "product barcode di temukan", product)
}

func (h *ProductHandler) Search(c *fiber.Ctx) error {
	products, err := h.service.Search(c.Query("q"))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "product tidak di temukan", err)
	}

	return response.Success(c, fiber.StatusOK, "berhasil mendapatkan data", products)
}

func (h *ProductHandler) SaveProduct(c *fiber.Ctx) error {
	var product *domain.Product
	err := c.BodyParser(&product)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "bad request", err.Error())
	}

	result, err := h.service.SaveProduct(product)
	if err != nil {
		if err.Error() == "sku already exists" {
			return response.Error(c, fiber.StatusBadRequest, "sku already exists", err.Error())
		}
		return response.Error(c, fiber.StatusInternalServerError, "failed to save product", err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "product berhasil disimpan", result)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	return response.Success(c, fiber.StatusOK, "product berhasil di update", nil)
}
