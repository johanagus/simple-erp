package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/johanagus/simple-erp/pkg/response"
)

func RequirePermission(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles := c.Locals("roles")
		if roles == nil {
			return response.Error(c, fiber.StatusForbidden, "Access denied", errors.New("roles not found"))
		}

		// Check if user has the required permission
		if !hasPermission(c, permission) {
			return response.Error(c, fiber.StatusForbidden, "Access denied", errors.New("Access forbidden: insufficient permissions"))
		}

		return c.Next()
	}
}

func hasPermission(c *fiber.Ctx, permission string) bool {
	roles, ok := c.Locals("roles").([]interface{})
	if !ok {
		return false
	}
	for _, r := range roles {
		if rs, ok := r.(string); ok && rs == permission {
			return true
		}
	}
	return false
}
