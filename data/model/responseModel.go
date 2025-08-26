package model

type Response struct {
	Status    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}