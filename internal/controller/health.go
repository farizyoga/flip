package controller

import (
	"flip/internal/helper/response"

	"github.com/gofiber/fiber/v2"
)

func (i *Controller) ActionHealthGet(c *fiber.Ctx) error {
	return response.CreateSuccess(c, nil)
}
