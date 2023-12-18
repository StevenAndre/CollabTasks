package entity

import (
	"time"
)

type PrivateMessage struct {
	PriMsgID   string    `gorm:"column:pm_id;primaryKey;size:10"`
	Message    string    `gorm:"column:message;not null"`
	SenderID   string    `gorm:"column:sender_id;not null;size:4"`
	ReceiverID string    `gorm:"column:receiver_id;not null;size:4"`
	IsRead     bool      `gorm:"column:isRead;default:false"`
	SentDate   time.Time `gorm:"column:sent_date;not null"`
	Sender     User   	 `gorm:"ForeignKey:SenderID"`
	Receiver   User		 `gorm:"ForeignKey:ReceiverID"`
}