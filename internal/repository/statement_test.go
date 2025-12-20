package repository_test

import (
	"context"
	"flip/internal/entity"
	"flip/internal/repository"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestNewStatementRepository(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want repository.StatementRepositoryMethod
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := repository.NewStatementRepository()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewStatementRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatementRepository_Create(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewStatementRepository()
	uploadID := "test"

	err := repo.Create(ctx, uploadID, entity.Statement{
		UploadID:     uploadID,
		ID:           uuid.NewString(),
		Timestamp:    1232522321,
		Counterparty: "Test",
		Type:         "CREDIT",
		Amount:       1000,
		Status:       "SUCCESS",
		Description:  "Test",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	statements, err := repo.Get(ctx, uploadID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(statements) != 1 {
		t.Fatalf("record should be created")
	}
}

func TestStatementRepository_Get(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewStatementRepository()
	uploadID := "test"

	statements, err := repo.Get(ctx, uploadID)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if len(statements) != 0 {
		t.Fatal("number of statement should 0")
	}

	err = repo.Create(ctx, uploadID, entity.Statement{
		UploadID:     uploadID,
		ID:           uuid.NewString(),
		Timestamp:    1232522321,
		Counterparty: "Test",
		Type:         "CREDIT",
		Amount:       1000,
		Status:       "SUCCESS",
		Description:  "Test",
	})
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	statements, err = repo.Get(ctx, uploadID)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if len(statements) != 1 {
		t.Fatal("number of statement should 1")
	}
}

func TestStatementRepository_GetWithPagination(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewStatementRepository()
	uploadID := "test"

	statements, _, err := repo.GetWithPagination(ctx, uploadID, repository.StatementFilter{
		Status: []string{"SUCCESS"},
	}, 1, 10)
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	if len(statements) != 0 {
		t.Fatal("number of data should 0")
	}

	err = repo.Create(ctx, uploadID, entity.Statement{
		UploadID:     uploadID,
		ID:           uuid.NewString(),
		Timestamp:    1232522321,
		Counterparty: "Test",
		Type:         "CREDIT",
		Amount:       1000,
		Status:       "SUCCESS",
		Description:  "Test",
	})
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	statements, _, err = repo.GetWithPagination(ctx, uploadID, repository.StatementFilter{
		Status: []string{"SUCCESS"},
		Type:   "CREDIT",
	}, 1, 10)
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	if len(statements) != 1 {
		t.Fatal("number of data should 1 based on filter")
	}
}

func TestStatementRepository_UpdateToSuccess(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewStatementRepository()
	uploadID := "test"
	id := uuid.NewString()

	err := repo.Create(ctx, uploadID, entity.Statement{
		UploadID:     uploadID,
		ID:           id,
		Timestamp:    1232522321,
		Counterparty: "Test",
		Type:         "CREDIT",
		Amount:       1000,
		Status:       "FAILED",
		Description:  "Test",
	})
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	err = repo.UpdateToSuccess(ctx, uploadID, id)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	statements, err := repo.Get(ctx, uploadID)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !strings.EqualFold(statements[0].Status, "SUCCESS") {
		t.Fatal("statement status should be updated to success")
	}
}

func TestStatementRepository_UpdateToFailed(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewStatementRepository()
	uploadID := "test"
	id := uuid.NewString()

	err := repo.Create(ctx, uploadID, entity.Statement{
		UploadID:     uploadID,
		ID:           id,
		Timestamp:    1232522321,
		Counterparty: "Test",
		Type:         "CREDIT",
		Amount:       1000,
		Status:       "SUCCESS",
		Description:  "Test",
	})
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	err = repo.UpdateToFailed(ctx, uploadID, id)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	statements, err := repo.Get(ctx, uploadID)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !strings.EqualFold(statements[0].Status, "FAILED") {
		t.Fatal("statement status should be updated to success")
	}
}
