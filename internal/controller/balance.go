package controller

import (
	"flip/internal/helper/response"

	"github.com/gofiber/fiber/v2"
)

func (i *Controller) ActionBalanceGet(c *fiber.Ctx) error {
	params := c.Queries()
	id := params["upload_id"]

	if id == "" {
		return response.CreateError(c, "upload_id is missing")
	}

	balance := i.Usecase.GetBalanceByUploadID(c.UserContext(), id)

	return response.CreateSuccess(c, map[string]interface{}{
		"balance": balance,
	})
}
