package model

import (
	"time"
)
type TimestampModel struct {
	BaseModel
	UserID   string    `json:"user_id" gorm:"type:uuid;foreignKey:UserID;references:ID"`
	CheckIn  time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
}
