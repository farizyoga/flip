package repository_test

import (
	"context"
	"flip/internal/entity"
	"flip/internal/repository"
	"testing"
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
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		uploadID string
		data     *entity.Statement
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var i repository.StatementRepository
			gotErr := i.Create(context.Background(), tt.uploadID, tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Create() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Create() succeeded unexpectedly")
			}
		})
	}
}

func TestStatementRepository_Get(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		uploadID string
		want     []*entity.Statement
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var i repository.StatementRepository
			got, gotErr := i.Get(context.Background(), tt.uploadID)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Get() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Get() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatementRepository_GetWithPagination(t *testing.T) {
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
			var i repository.StatementRepository
			got, gotErr := i.GetWithPagination(context.Background(), tt.uploadID, tt.filter, tt.page, tt.size)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetWithPagination() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetWithPagination() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetWithPagination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatementRepository_UpdateToSuccess(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		uploadID string
		id       string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var i repository.StatementRepository
			gotErr := i.UpdateToSuccess(context.Background(), tt.uploadID, tt.id)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateToSuccess() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateToSuccess() succeeded unexpectedly")
			}
		})
	}
}

func TestStatementRepository_UpdateToFailed(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		uploadID string
		id       string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var i repository.StatementRepository
			gotErr := i.UpdateToFailed(context.Background(), tt.uploadID, tt.id)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateToFailed() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateToFailed() succeeded unexpectedly")
			}
		})
	}
}
