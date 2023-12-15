package entity

type Position struct {
	PositionID          int16  `gorm:"column:id_position;primaryKey"`
	Name        string `gorm:"column:name;not null;size:50"`
	Description string `gorm:"column:description;size:100"`
	TeamMembers   []TeamMember	 `gorm:"ForeignKey:PositionID"`
}

func (p *Position) TableName() string {
	return "positions"
}