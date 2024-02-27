package model

import "time"

type Evidence struct {
	EvidenceID string    `json:"evidence_id"`
	FilePath   string    `json:"filepath"`
	Comment    string    `json:"comment"`
	TaskID     string    `json:"task_id"`
	CreatedAt  time.Time `json:"create_at"`
}
