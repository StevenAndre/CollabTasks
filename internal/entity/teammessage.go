package entity

import "time"

type TeamMessage struct {
	TeamMsgID   string    `gorm:"column:tm_id;primaryKey;size:10"`
	TeamID   	string    `gorm:"column:team_id;not null;size:8"`
	MemberID 	string    `gorm:"column:member_id;not null;size:4"`
	Message  	string    `gorm:"column:message;not null"`
	SendDate 	time.Time `gorm:"column:send_date;not null"`
	Sender      User  	  `gorm:"ForeignKey:MemberID"`	
	Team    	Team	  `gorm:"ForeignKey:TeamID"`
}