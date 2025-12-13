package model

type Menu struct {
	BaseModel
	Name		string  `json:"name"`
	Price		float64 `json:"price"`
	Description	string  `json:"description"`
	Items      []Item   `json:"items" gorm:"many2many:menu_items;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;" swaggerignore:"true"`
}