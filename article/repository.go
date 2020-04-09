package article

import (
	"context"

	"github.com/dynastymasra/whistleblower/domain"
)

// TODO: For demo only and more interface in the future
type Repository interface {
	Create(context.Context, domain.Article) error
	FindAll(context.Context, map[string]interface{}) ([]*domain.Article, error)
	Delete(context.Context, domain.Article) error
}
