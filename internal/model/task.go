package model

import (
	"time"
)

type Task struct {
	TaskID      string    `json:"task_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	IsCompleted bool      `json:"is_completed"`
	UserID      string    `json:"user_id"`
	ProjectID   string    `json:"project_id"`
}

type TaskU struct {
	TaskID      string    `json:"task_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	IsCompleted bool      `json:"is_completed"`
	UserID      string    `json:"user_id"`
	ProjectID   string    `json:"project_id"`
	User        User      `json:"user"`
}
