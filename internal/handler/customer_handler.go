package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/service"
	"github.com/johanagus/simple-erp/pkg/response"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) FindAll(c *fiber.Ctx) error {
	customers, err := h.service.FindAllCustomer()
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "customer tidak di temukan", err)
	}

	return response.Success(c, fiber.StatusOK, "berhasil mendapatkan data", customers)
}

func (h *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	var customer *domain.Customer
	err := c.BodyParser(&customer)

	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "bad request", err.Error())
	}

	err = h.service.CreateCustomer(customer)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "failed to save customer", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "customer berhasil disimpan", nil)
}

func (h *CustomerHandler) FindByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid customer ID", err)
	}

	customer, err := h.service.FindByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "customer id : "+strconv.Itoa(id)+" is not found", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "customer id : "+strconv.Itoa(id)+" found", customer)
}

func (h *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid customer ID", err)
	}

	customer := &domain.Customer{}
	if err := c.BodyParser(customer); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", err)
	}

	err = h.service.UpdateCustomer(id, customer)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "failed to update customer", err)
	}

	return response.Success(c, fiber.StatusOK, "customer berhasil di update", nil)
}
