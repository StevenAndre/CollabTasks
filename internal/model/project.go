package model

import "time"

type Project struct {
	ProjectID   string    `json:"project_id"`
	Name       string      `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
	CreatedBy   string    `json:"created_by"`
}