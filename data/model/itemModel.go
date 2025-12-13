package model

type Item struct {
	BaseModel
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Menu        []Menu   `json:"menus" gorm:"many2many:menu_items;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;" swaggerignore:"true"`
}
