package app

import (
	"context"
	"flip/internal/consumer"
	"flip/internal/controller"
	"flip/internal/repository"
	"flip/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(ctx context.Context) (*fiber.App, usecase.UsecaseMethod) {
	statementRepository := repository.NewStatementRepository()
	reconciliationConsumer := consumer.NewReconciliationConsumer(5, 100, 5, statementRepository)

	go reconciliationConsumer.Listen(ctx)

	uc := usecase.NewUsecase(usecase.Usecase{
		StatementRepository:    statementRepository,
		ReconciliationConsumer: reconciliationConsumer,
	})

	ctrl := controller.NewController(controller.Controller{
		Usecase: uc,
	})

	r := fiber.New()
	r.Use(logger.New())

	r.Post("/statements", ctrl.ActionStatementPost)
	r.Get("/balance", ctrl.ActionBalanceGet)
	r.Get("/transactions/issues", ctrl.ActionTransactionIssueGet)
	r.Get("/health", ctrl.ActionHealthGet)

	return r, uc
}
