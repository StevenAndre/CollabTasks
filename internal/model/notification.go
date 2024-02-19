package model

import "time"

type Notification struct {
	NotificationID string 		`json:"Notification_id"`
	Issue          string 		`json:"issue"`
	Message        string 	`    json:"message"`
	Issued         time.Time 	`json:"issued"`
	Isread         bool 		`json:"isread"`
	UserID         string  		`json:"user_id"`
}