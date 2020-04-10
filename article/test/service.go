package test

import (
	"context"

	"github.com/dynastymasra/whistleblower/domain"
	"github.com/stretchr/testify/mock"
)

type MockArticleService struct {
	mock.Mock
}

func (m *MockArticleService) Create(ctx context.Context, article domain.Article) (*domain.Article, error) {
	args := m.Called(ctx, article)
	return args.Get(0).(*domain.Article), args.Error(1)
}
