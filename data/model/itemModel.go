package model

type Item struct {
	BaseModel
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price      float64   `json:"price"`
}
