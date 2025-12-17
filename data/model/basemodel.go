package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string   `json:"id" gorm:"type:uuid;primaryKey" swaggerignore:"true"`		
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggerignore:"true"`
}