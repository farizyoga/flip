package usecase_test

import (
	"context"
	"flip/internal/consumer"
	"flip/internal/entity"
	"flip/internal/repository"
	"flip/internal/usecase"
	"testing"
)

func TestNewUsecase(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		opt  usecase.Usecase
		want usecase.UsecaseMethod
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := usecase.NewUsecase(tt.opt)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_CreateStatement(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		id      string
		data    *entity.Statement
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var i usecase.Usecase
			gotErr := i.CreateStatement(context.Background(), tt.id, tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateStatement() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateStatement() succeeded unexpectedly")
			}
		})
	}
}

func TestUsecase_GetBalanceByUploadID(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		id   string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var i usecase.Usecase
			got := i.GetBalanceByUploadID(context.Background(), tt.id)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetBalanceByUploadID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_FindStatementByUploadID(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		uploadID string
		filter   repository.StatementFilter
		page     int
		size     int
		want     []*entity.Statement
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var i usecase.Usecase
			got, gotErr := i.FindStatementByUploadID(context.Background(), tt.uploadID, tt.filter, tt.page, tt.size)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("FindStatementByUploadID() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("FindStatementByUploadID() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("FindStatementByUploadID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_Publish(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		message consumer.ReconciliationConsumerMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var i usecase.Usecase
			i.Publish(context.Background(), tt.message)
		})
	}
}

func TestUsecase_UpdateAllStatementToFailed(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		uploadID string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var i usecase.Usecase
			gotErr := i.UpdateAllStatementToFailed(context.Background(), tt.uploadID)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateAllStatementToFailed() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateAllStatementToFailed() succeeded unexpectedly")
			}
		})
	}
}
