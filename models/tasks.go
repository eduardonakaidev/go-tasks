package models

import "time"

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
	IdUser      string `json:"id_user"`
	CreatedAt   time.Time `json:"created_at"`
}

// representa o input para criar a task
type TaskInputPublic struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type TaskInputUpdate struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}