package article

import (
	"context"

	"github.com/dynastymasra/cookbook"

	"github.com/dynastymasra/whistleblower/domain"
)

type Service interface {
	Create(context.Context, domain.Article) (*domain.Article, error)
}

type ServiceInstance struct {
	repo Repository
}

func NewService(repo Repository) ServiceInstance {
	return ServiceInstance{repo: repo}
}

func (s ServiceInstance) Create(ctx context.Context, article domain.Article) (*domain.Article, error) {
	if err := s.repo.Create(ctx, article); err != nil {
		return nil, err
	}

	// Get user if success insert to database, return the result to client
	res, err := s.repo.Find(ctx, cookbook.JSON{"id": article.ID})
	if err != nil {
		return &article, nil
	}

	return res, nil
}
