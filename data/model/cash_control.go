package model

type CashControl struct {
	BaseModel
	Cash float64 `json:"cash"`
	Coin float64 `json:"coin"`
	BreadIn float64 `json:"bread_in"`
	Safe float64 `json:"safe"`
	UserID string `json:"user_id" gorm:"type:uuid;foreignKey:UserID;references:ID"`
}