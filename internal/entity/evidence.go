package entity

import "time"

type Evidence struct {
	EvidenceID        string    `gorm:"column:id_evidence;primaryKey;size:10"`
	FilePath  string    `gorm:"column:filepath;not null;size:300"`
	Comment   string    `gorm:"column:comment;size:200"`
	TaskID    string    `gorm:"column:task_id;not null;size:10"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	Task Task
}