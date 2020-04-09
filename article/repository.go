package article

import (
	"context"

	"github.com/dynastymasra/whistleblower/config"
	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/whistleblower/domain"
)

// TODO: For demo only and more interface in the future
type Repository interface {
	Create(context.Context, domain.Article) error
	Find(context.Context, map[string]interface{}) (*domain.Article, error)
	FindAll(context.Context, map[string]interface{}) ([]*domain.Article, error)
}

type RepositoryInstance struct {
	db        *gorm.DB
	TableName string
}

func NewRepository(db *gorm.DB) *RepositoryInstance {
	return &RepositoryInstance{
		db:        db,
		TableName: config.ArticleTableName,
	}
}

// TODO: Add instrumentation to monitor performance (Newrelic, StatsD)
func (a RepositoryInstance) Create(ctx context.Context, article domain.Article) error {
	return a.db.Table(a.TableName).Omit(config.CreatedAtFieldName, config.UpdatedAtFieldName).Create(&article).Error
}

func (a RepositoryInstance) Find(ctx context.Context, filter map[string]interface{}) (*domain.Article, error) {
	var result *domain.Article

	if err := a.db.Table(a.TableName).First(&result, filter).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a RepositoryInstance) FindAll(ctx context.Context, filter map[string]interface{}) ([]*domain.Article, error) {
	var result []*domain.Article

	if err := a.db.Table(a.TableName).Where(filter).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
