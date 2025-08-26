package model

import (
	"time"
)

type Delivery struct {
	BaseModel
	OrderID   string    `json:"order_id"`
	ItemID    []string  `json:"item_id" gorm:"type:text[]"`
	Quantity  []int     `json:"quantity" gorm:"type:int[]"`
	Delivered bool      `json:"delivered"`
	DeliveredAt time.Time `json:"delivered_at"`
	UserID     string    `json:"user_id"`
	User   User `json:"user" gorm:"constraint:OnDelete:CASCADE;OnUpdate:CASCADE;foreignKey:UserID;references:ID" swaggerignore:"true"`
}