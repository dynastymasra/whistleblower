package test

import (
	"context"

	"github.com/dynastymasra/whistleblower/domain"
	"github.com/dynastymasra/whistleblower/infrastructure/provider"
	"github.com/stretchr/testify/mock"
)

type MockViewerRepository struct {
	mock.Mock
}

func (m *MockViewerRepository) Create(ctx context.Context, viewer domain.Viewer) error {
	args := m.Called(ctx, viewer)
	return args.Error(0)
}

func (m *MockViewerRepository) FindAll(ctx context.Context, query *provider.Query) ([]*domain.Viewer, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*domain.Viewer), args.Error(1)
}
