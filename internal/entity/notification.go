package entity

import (
	"time"

	
)

type Notification struct {
	NotificationID      string    `gorm:"primaryKey;column:id_notification"`
	Issue   string    `gorm:"not null;column:issue"`
	Message string    `gorm:"not null;column:message"`
	Issued  time.Time `gorm:"not null;column:issued"`
	Isread  bool      `gorm:"not null;column:is_read"`
	User 	User
	UserID  string    `gorm:"not null;column:id_user"`
}

func (n *Notification) TableName() string {
	return "notifications"
}