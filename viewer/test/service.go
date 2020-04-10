package test

import (
	"context"

	"github.com/dynastymasra/cookbook"
	"github.com/dynastymasra/whistleblower/infrastructure/provider"
	"github.com/stretchr/testify/mock"
)

type MockViewerService struct {
	mock.Mock
}

func (m *MockViewerService) Statistic(ctx context.Context, query *provider.Query) (*cookbook.JSON, error) {
	args := m.Called(ctx, query)
	return args.Get(0).(*cookbook.JSON), args.Error(1)
}
