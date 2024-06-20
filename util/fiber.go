package util

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(code int, msg string) fiber.Map {
	return fiber.Map{
		"success": false,
		"code":    code,
		"msg":     msg,
	}
}

func SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		"success": true,
		"code":    200,
		"data":    data,
	}
}
