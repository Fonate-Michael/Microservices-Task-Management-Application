package model

type Task struct {
	Id          int    `json:"id"`
	User_id     int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
