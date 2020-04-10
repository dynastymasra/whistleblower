package test

import (
	"context"

	"github.com/dynastymasra/whistleblower/domain"
	"github.com/stretchr/testify/mock"
)

type MockArticleRepository struct {
	mock.Mock
}

func (m *MockArticleRepository) Create(ctx context.Context, article domain.Article) error {
	args := m.Called(ctx, article)
	return args.Error(0)
}

func (m *MockArticleRepository) Find(ctx context.Context, update map[string]interface{}) (*domain.Article, error) {
	args := m.Called(ctx, update)
	return args.Get(0).(*domain.Article), args.Error(1)
}

func (m *MockArticleRepository) FindAll(ctx context.Context, query map[string]interface{}) ([]*domain.Article, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*domain.Article), args.Error(1)
}
