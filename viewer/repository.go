package viewer

import (
	"context"

	"github.com/dynastymasra/whistleblower/infrastructure/provider"

	"github.com/dynastymasra/whistleblower/config"
	"github.com/dynastymasra/whistleblower/domain"
	"github.com/jinzhu/gorm"
)

// TODO: For demo only and more interface in the future
type Repository interface {
	Create(context.Context, domain.Viewer) error
	FindAll(context.Context, *provider.Query) ([]*domain.Viewer, error)
}

type RepositoryInstance struct {
	db        *gorm.DB
	TableName string
}

func NewRepository(db *gorm.DB) *RepositoryInstance {
	return &RepositoryInstance{
		db:        db,
		TableName: config.ViewerTableName,
	}
}

// TODO: Add instrumentation to monitor performance (Newrelic, StatsD)
func (r RepositoryInstance) Create(ctx context.Context, viewer domain.Viewer) error {
	return r.db.Table(r.TableName).Create(&viewer).Error
}

func (r RepositoryInstance) FindAll(ctx context.Context, filter *provider.Query) ([]*domain.Viewer, error) {
	var result []*domain.Viewer

	db := r.db.Table(r.TableName)
	db = provider.TranslateQuery(db, filter)

	err := db.Find(&result).Error

	return result, err
}
