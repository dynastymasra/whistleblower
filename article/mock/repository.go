package mock

import (
	"context"

	"github.com/dynastymasra/whistleblower/domain"
	"github.com/stretchr/testify/mock"
)

type ArticleRepository struct {
	mock.Mock
}

func (m *ArticleRepository) Create(ctx context.Context, article domain.Article) error {
	args := m.Called(ctx, article)
	return args.Error(0)
}

func (m *ArticleRepository) Find(ctx context.Context, update map[string]interface{}) (*domain.Article, error) {
	args := m.Called(ctx, update)
	return args.Get(0).(*domain.Article), args.Error(1)
}

func (m *ArticleRepository) FindAll(ctx context.Context, query map[string]interface{}) ([]*domain.Article, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*domain.Article), args.Error(1)
}
