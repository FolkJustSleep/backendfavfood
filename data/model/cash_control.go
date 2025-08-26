package model

type CashControl struct {
	BaseModel
	UserID string  `json:"user_id"`
	Cash float64 `json:"cash"`
	Coin float64 `json:"coin"`
	BreadIn float64 `json:"bread_in"`
	Safe float64 `json:"safe"`
	User   User `json:"user" gorm:"constraint:OnDelete:CASCADE;OnUpdate:CASCADE;foreignKey:UserID;references:ID" swaggerignore:"true"`
}