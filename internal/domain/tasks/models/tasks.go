package models

type Task struct {
	ID          int    `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
