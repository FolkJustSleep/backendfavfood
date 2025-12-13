package model

import( 
	"time"
)

type User struct {
	BaseModel
	Name	string `json:"name"`
	Email	string `json:"email" gorm:"unique"`
	Password	string `json:"password" gorm:"not null"`
	Role	string `json:"role" gorm:"default:'user'"`
	CashControl []CashControl `json:"cash_controls" gorm:"many2many:cash_control_users;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;" swaggerignore:"true"`
}

type Timestamp struct {
	BaseModel
	UserID string `json:"user_id"`
	In     time.Time  `json:"in"`
	Out    time.Time  `json:"out"`
	User   User `json:"user" gorm:"constraint:OnDelete:CASCADE;OnUpdate:CASCADE;foreignKey:UserID;references:ID" swaggerignore:"true"`
}