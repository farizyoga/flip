package controller

import (
	"flip/internal/helper/response"
	"flip/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func (i *Controller) ActionTransactionIssueGet(c *fiber.Ctx) error {
	uploadID := c.Query("upload_id")
	if uploadID == "" {
		return response.CreateFail(c, "upload_id is missing")
	}

	page := c.QueryInt("page", 1)
	size := c.QueryInt("size", 10)

	statements, err := i.Usecase.FindStatementByUploadID(c.UserContext(), uploadID, repository.StatementFilter{
		Status: "FAILED",
		Type:   "CREDIT",
	}, page, size)
	if err != nil {
		return response.CreateError(c, err.Error())
	}

	return response.CreateSuccess(c, map[string]interface{}{
		"rows": statements,
		"page": page,
	})
}
