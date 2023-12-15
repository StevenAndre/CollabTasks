package entity

type AssignedProject struct {
	ProjectID string `gorm:"column:id_project;primaryKey;size:8"`
	TeamID    string `gorm:"column:id_team;primaryKey;size:8"`
	Project  Project `gorm:"ForeignKey:ProjectID"`
	Team     Team   `gorm:"ForeignKey:TeamID"`
}

func (ap *AssignedProject) TableName() string {
	return "assigned_project"
}
