package entity

import "time"

type Task struct {
	TaskID      string    `gorm:"column:id_task;primaryKey;size:10"`
	Title       string    `gorm:"column:title;not null;size:100"`
	Description string    `gorm:"column:description;not null;size:250"`
	CreatedAt   time.Time `gorm:"column:created_at;not null"`
	IsCompleted bool      `gorm:"column:iscompleted;default:false"`
	UserID      string    `gorm:"column:id_user;not null;size:4"`
	ProjectID   string    `gorm:"column:id_project;not null;size:8"`
	Evidences   []Evidence `gorm:"ForeignKey:TaskID"`
	User 	User
	Project Project

}