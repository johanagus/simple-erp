package handler

import (
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/dto"
	"github.com/johanagus/simple-erp/internal/service"
	"github.com/johanagus/simple-erp/pkg/response"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{service: userService}
}

func (h *UserHandler) FindAll(c *fiber.Ctx) error {
	user, err := h.service.FindAll()
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "user tidak di temukan", err)
	}

	result := make([]dto.UserRespons, len(user))
	for i, u := range user {
		result[i] = dto.ToUserRespons(&u)
	}

	return response.Success(c, fiber.StatusOK, "berhasil mendapatkan data", result)

}

func (h *UserHandler) FindByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) // tangkap parameter id dan konversi dari string ke int
	user, err := h.service.FindByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "user id : "+strconv.Itoa(id)+" is not found", err.Error())
	}

	result := dto.ToUserRespons(&user)

	return response.Success(c, fiber.StatusOK, "user id di temukan", result)
}

func (h *UserHandler) SaveUser(c *fiber.Ctx) error {
	user := &domain.User{}

	// Parse request body into user struct
	if err := c.BodyParser(user); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "bad request", err.Error())
	}

	// Validate user struct
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "validation failed", err.Error())
	}

	// Save user
	if err := h.service.SaveUser(user); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "failed to save user", err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "user berhasil disimpan", nil)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := &domain.User{}
	if err := c.BodyParser(user); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", err)
	}

	err := h.service.UpdateUser(id, user)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "failed to update user", err)
	}

	return response.Success(c, fiber.StatusOK, "user berhasil di update", nil)
}
