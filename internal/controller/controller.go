package controller

import (
	"flip/internal/consumer"
	"flip/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Usecase                usecase.UsecaseMethod
	ReconciliationConsumer consumer.ReconciliationConsumerMethod
}

func NewController(opt Controller) *Controller {
	return &opt
}

func (i *Controller) ActionIndex(c *fiber.Ctx) error {
	return c.SendFile("html/index.html")
}
