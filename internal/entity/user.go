package entity

import (

	
)

//TODO mapear para asegurar todos los campos con "column" como el id aqui
type User struct {
    
	UserID        string         `gorm:"column:id_user;primaryKey;size:4"`
	Name          string         `gorm:"column:name;not null;size:60"`
	Lastname      string         `gorm:"column:lastname;not null;size:70"`
	Username      string         `gorm:"column:username;not null;unique;size:60"`
	Email         string         `gorm:"column:email;not null;unique;size:200"`
	Password      string         `gorm:"column:password;not null;size:255"`
	ProfilePhoto  string         `gorm:"column:profile_photo;size:300"`
	Notifications []Notification `gorm:"ForeignKey:UserID"`
}

func (u *User) TableName() string {
	return "users"
}
