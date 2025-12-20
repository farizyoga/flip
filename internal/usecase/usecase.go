package usecase

import (
	"context"
	"flip/internal/consumer"
	"flip/internal/entity"
	"flip/internal/repository"
)

type Usecase struct {
	StatementRepository    repository.StatementRepositoryMethod
	ReconciliationConsumer consumer.ReconciliationConsumerMethod
}

type UsecaseMethod interface {
	CreateStatement(ctx context.Context, id string, data *entity.Statement) error
	GetBalanceByUploadID(ctx context.Context, id string) int
	FindStatementByUploadID(ctx context.Context, uploadID string, filter repository.StatementFilter, page, size int) ([]*entity.Statement, int, error)
	Publish(ctx context.Context, message consumer.ReconciliationConsumerMessage)
	UpdateAllStatementToFailed(ctx context.Context, uploadID string) error
}

func NewUsecase(opt Usecase) UsecaseMethod {
	return &opt
}

func (i *Usecase) CreateStatement(ctx context.Context, id string, data *entity.Statement) error {
	return i.StatementRepository.Create(ctx, id, data)
}

func (i *Usecase) GetBalanceByUploadID(ctx context.Context, id string) int {
	statements, err := i.StatementRepository.Get(ctx, id)

	if err != nil {
		return 0
	}
	balance := 0
	for _, s := range statements {
		if s.IsFailed() {
			continue
		}

		if s.IsCredit() {
			balance += s.Amount
		}

		if s.IsDebit() {
			balance -= s.Amount
		}
	}

	return balance
}

func (i Usecase) FindStatementByUploadID(ctx context.Context, uploadID string, filter repository.StatementFilter, page, size int) ([]*entity.Statement, int, error) {
	return i.StatementRepository.GetWithPagination(ctx, uploadID, filter, page, size)
}

func (i *Usecase) Publish(ctx context.Context, message consumer.ReconciliationConsumerMessage) {
	i.ReconciliationConsumer.Publish(ctx, message)
}

func (i *Usecase) UpdateAllStatementToFailed(ctx context.Context, uploadID string) error {
	statements, _ := i.StatementRepository.Get(ctx, uploadID)
	for _, s := range statements {
		i.StatementRepository.UpdateToFailed(ctx, s.UploadID, s.ID)
	}

	return nil
}
