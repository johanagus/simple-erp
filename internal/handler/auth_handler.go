package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/internal/service"
	"github.com/johanagus/simple-erp/pkg/response"
	"github.com/johanagus/simple-erp/pkg/utils"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Signin(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Input tidak valid", err.Error())
	}

	user, err := h.service.Authenticate(input.Email, input.Password)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Authentikasi gagal", err.Error())
	}

	token, err := utils.GenerateAccessToken(int(user.ID), user.Email, user.Roles)
	if err != nil {
		return response.Error(c, 500, "Gagal membuat token", err.Error())
	}

	refreshToken, err := utils.GenerateRefreshToken(int(user.ID))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Gagal membuat refresh token", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Sign in berhasil", fiber.Map{
		"token":         token,
		"refresh_token": refreshToken,
		"user": fiber.Map{
			"id":        user.ID,
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
			"email":     user.Email,
			"roles":     user.Roles,
		},
	})
}
