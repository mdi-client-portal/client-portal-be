package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

func Success(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, code int, message string, errors interface{}) error {
	return c.Status(code).JSON(Response{
		Code:    code,
		Message: message,
		Errors:  errors,
	})
}