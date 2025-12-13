package model

type Logs struct {
	BaseModel
	UserID string `json:"user_id" gorm:"type:uuid;foreignKey:UserID;references:ID"`
	Action  string `json:"action"`
	Status  string `json:"status"`
}