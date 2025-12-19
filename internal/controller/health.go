package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (i *Controller) ActionHealthGet(c *fiber.Ctx) error {
	return c.JSON(i.ReconciliationConsumer.GetStatus(c.UserContext()))
}
