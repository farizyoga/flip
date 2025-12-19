package repository

import (
	"context"
	"flip/internal/entity"

	"sort"
	"strings"
	"sync"
)

type StatementRepository struct {
	mu   sync.RWMutex
	data map[string]map[string]*entity.Statement
}

type StatementFilter struct {
	Status string
	Type   string
}

type StatementRepositoryMethod interface {
	Create(ctx context.Context, uploadID string, data *entity.Statement) error
	Get(ctx context.Context, uploadID string) ([]*entity.Statement, error)
	GetWithPagination(ctx context.Context, uploadID string, filter StatementFilter, page, size int) ([]*entity.Statement, error)
	UpdateToSuccess(ctx context.Context, uploadID string, id string) error
	UpdateToFailed(ctx context.Context, uploadID string, id string) error
}

func NewStatementRepository() StatementRepositoryMethod {
	return &StatementRepository{
		mu:   sync.RWMutex{},
		data: make(map[string]map[string]*entity.Statement),
	}
}

func (i *StatementRepository) Create(ctx context.Context, uploadID string, data *entity.Statement) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	_, ok := i.data[uploadID]
	if !ok {
		i.data[uploadID] = map[string]*entity.Statement{}
	}

	i.data[uploadID][data.ID] = data

	return nil
}

func (i *StatementRepository) Get(ctx context.Context, uploadID string) ([]*entity.Statement, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	data := make([]*entity.Statement, 0)
	statements, ok := i.data[uploadID]
	if !ok {
		return data, nil
	}

	for _, s := range statements {
		data = append(data, s)
	}

	return data, nil
}

func (i *StatementRepository) GetWithPagination(ctx context.Context, uploadID string, filter StatementFilter, page, size int) ([]*entity.Statement, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	data := i.data[uploadID]
	filtered := []*entity.Statement{}

	for _, s := range data {
		if filter.Status != "" {
			if !strings.EqualFold(s.Status, filter.Status) {
				continue
			}
		}

		if filter.Type != "" {
			if !strings.EqualFold(s.Type, filter.Type) {
				continue
			}
		}

		filtered = append(filtered, s)
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].ID > filtered[j].ID
	})

	start := (page - 1) * size

	if start >= len(filtered) {
		return make([]*entity.Statement, 0), nil
	}

	end := start + size
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], nil
}

func (i *StatementRepository) UpdateToSuccess(ctx context.Context, uploadID string, id string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.data[uploadID][id].Status = "SUCCESS"
	return nil
}

func (i *StatementRepository) UpdateToFailed(ctx context.Context, uploadID string, id string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.data[uploadID][id].Status = "FAILED"
	return nil
}
