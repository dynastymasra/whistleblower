package domain

import "time"

type (
	// TODO: Add more field this struct for demo only
	Article struct {
		ID        string     `json:"id" validate:"required" gorm:"column:id"`
		CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
		UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	}

	Viewer struct {
		ID        string     `json:"id" validate:"required" gorm:"column:id"`
		ArticleID string     `json:"id" validate:"required" gorm:"column:article_id"`
		CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	}
)
