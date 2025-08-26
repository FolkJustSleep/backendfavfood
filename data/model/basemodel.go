package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string   `json:"id" gorm:"type:uuid;primaryKey"`		
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggerignore:"true"`
}