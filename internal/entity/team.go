package entity

type Team struct {
	TeamID   string `gorm:"column:id_team;primaryKey;size:8"`
	Name string `gorm:"column:name;not null;size:100"`
	TeamMembers   []TeamMember	 `gorm:"ForeignKey:TeamID"`
	AssignedProjects[]AssignedProject  `gorm:"ForeignKey:TeamID"`

}

func (t *Team) TableName() string {
	return "teams"
}
