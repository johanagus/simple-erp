package response

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

func Success(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})

}

func Error(c *fiber.Ctx, code int, message string, errors interface{}) error {
	return c.Status(code).JSON(ErrorResponse{
		Code:    code,
		Message: message,
		Errors:  errors,
	})
}
