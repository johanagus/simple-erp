package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/service"
	"github.com/johanagus/simple-erp/pkg/response"
)

type StoreHandler struct {
	service service.StoreService
}

func NewStoreHandler(service service.StoreService) *StoreHandler {
	return &StoreHandler{service: service}
}

func (h *StoreHandler) FindAll(c *fiber.Ctx) error {
	store, err := h.service.GetAll()
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "store tidak di temukan", err)
	}

	if len(*store) == 0 {
		return response.Error(c, fiber.StatusNotFound, "store tidak di temukan", nil)
	}

	return response.Success(c, fiber.StatusOK, "berhasil mendapatkan data", store)
}

func (h *StoreHandler) FindByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	store, err := h.service.FindByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "store id : "+strconv.Itoa(id)+" is not found", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "store id di temukan", store)
}

func (h *StoreHandler) CreateStore(c *fiber.Ctx) error {
	store := &domain.Store{}
	if err := c.BodyParser(store); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", err)
	}

	_, err := h.service.Create(store)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "failed to create store", err)
	}

	return response.Success(c, fiber.StatusOK, "store berhasil di tambahkan", nil)
}

func (h *StoreHandler) UpdateStore(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	store := &domain.Store{}
	if err := c.BodyParser(store); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", err)
	}

	h.service.Update(id, store)
	return response.Success(c, fiber.StatusOK, "store berhasil di update", nil)
}
