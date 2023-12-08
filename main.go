package main

import (
	"context"
	"fmt"

	"github.com/StevenAndre/collabtasks/database"
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/settings"
)

func main() {

	s,err := settings.NewSettings()
	if err!=nil{
		panic(err)
	}
	fmt.Println(s)
	

	db,err:=database.NewDatabase(context.Background(),s)
	if err!=nil{
		panic(err)
	}

	
	var users []entity.User
	result := db.Preload("Notifications").Find(&users)
	if result.Error != nil {
		panic("Failed to fetch users")
	}

	for _, user := range users {
		fmt.Println("User ID:", user.UserID)
		fmt.Println("Name:", user.Name)
		fmt.Println("Notifications:")

		for _, notification := range user.Notifications {
			fmt.Println("  - Issue:", notification.Issue)
			fmt.Println("    Message:", notification.Message)
			fmt.Println("    Issued:", notification.Issued)
		}
	}



}
