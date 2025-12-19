package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateSuccess(c *fiber.Ctx, data interface{}) error {
	return c.JSON(map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}

func CreateError(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
		"status":  "error",
		"message": message,
	})
}

func CreateFail(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusUnprocessableEntity).JSON(map[string]interface{}{
		"status": "fail",
		"data":   data,
	})
}
