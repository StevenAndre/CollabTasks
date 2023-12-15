package entity

import "time"

type Project struct {
	ProjectID   string    `gorm:"column:id_project;primaryKey;size:8"`
	Name        string    `gorm:"column:name;not null;size:100"`
	Description string    `gorm:"column:description;size:250"`
	CreatedAt   time.Time `gorm:"column:created_at;not null"`
	User 	User		  `gorm:"foreignKey:CreatedBy"`
	CreatedBy   string    `gorm:"column:create_by;not null;size:4"`
	Tasks 		[]Task    `gorm:"ForeignKey:ProjectID"`
	AssignedProjects []AssignedProject  `gorm:"ForeignKey:ProjectID"`
	
}

func (p *Project) TableName() string {
	return "projects"
}