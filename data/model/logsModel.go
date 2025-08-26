package model

type Logs struct {
	BaseModel
	UserID string `json:"user_id"`
	Action  string `json:"action"`
	Status  string `json:"status"`
}