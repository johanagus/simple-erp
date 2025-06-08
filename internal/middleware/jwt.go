package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/johanagus/simple-erp/config"
	"github.com/johanagus/simple-erp/pkg/response"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return response.Error(c, fiber.StatusUnauthorized, "Token tidak ditemukan", nil)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetEnv("JWT_SECRET", "secret")), nil
		})

		if err != nil || !token.Valid {
			return response.Error(c, fiber.StatusUnauthorized, "Token tidak valid", err.Error())
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return response.Error(c, fiber.StatusUnauthorized, "Token tidak valid", nil)
		}

		// Simpan user info ke context, bisa dipakai di handler
		c.Locals("id", claims["id"])
		c.Locals("email", claims["email"])

		return c.Next()
	}
}
