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
func (r RepositoryInstance) Create(ctx context.Context, article domain.Article) error {
	return r.db.Table(r.TableName).Create(&article).Error
}

func (r RepositoryInstance) Find(ctx context.Context, query map[string]interface{}) (*domain.Article, error) {
	var result domain.Article

	if err := r.db.Table(r.TableName).First(&result, query).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r RepositoryInstance) FindAll(ctx context.Context, filter map[string]interface{}) ([]*domain.Article, error) {
	var result []*domain.Article

	err := r.db.Table(r.TableName).Where(filter).Find(&result).Error

	return result, err
}
