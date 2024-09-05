package entity

type TeamMember struct {
	UserID         string `gorm:"column:id_member;primaryKey;size:4"`
	TeamID     string `gorm:"column:id_team;primaryKey;size:8"`
	PositionID int16  `gorm:"column:position_id;not null"`
	User 	User	  `gorm:"foreignKey:UserID"`
	Team  	Team      `gorm:"foreignKey:TeamID"`
	Position Position `gorm:"foreignKey:PositionID"`
}

func (t *TeamMember) TableName() string {
	return "team_members"
}
