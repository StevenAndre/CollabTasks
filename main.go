package main

import (
	"context"
	"fmt"

	"github.com/StevenAndre/collabtasks/database"
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/repository"
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
	var projects []entity.Project
	
	


	projRs:=db.Find(&projects)
	result := db.Preload("Notifications").Preload("TeamMembers").Preload("Projects.Tasks.Evidences").Preload("Projects").Preload("Projects.AssignedProjects").Find(&users)

	

	if result.Error != nil {
		panic("Failed to fetch users")
	}

	if projRs.Error != nil {
		panic("Failed to fetch projects")
	}

	for _, user := range users {
		fmt.Println("User ID:", user.UserID)
		fmt.Println("Name:", user.Name)
		fmt.Println("Notifications:")

		for _, tm := range user.TeamMembers {
			fmt.Println("  - PositionID:", tm.PositionID)
			fmt.Println("    TeamID:", tm.TeamID)
			fmt.Println("    UserID:", tm.UserID)
		
		}



		for _, notification := range user.Notifications {
			fmt.Println("  - Issue:", notification.Issue)
			fmt.Println("    Message:", notification.Message)
			fmt.Println("    Issued:", notification.Issued)
		}


		
	fmt.Println("------------------------PROJECTS------------------------")
	for _, pr := range user.Projects {
		fmt.Println("PRoject ID:", pr.ProjectID)
		fmt.Println("Name:", pr.Name)
		fmt.Println("Description:", pr.Description)
		fmt.Println("USerID:", pr.CreatedBy)
		fmt.Println("TASKS:")

		for _, task := range pr.Tasks {
			fmt.Println("  DESCription :", task.Description)
			fmt.Println("  ProjecID :", task.ProjectID)
			fmt.Println(" Title   :", task.Title)
			fmt.Println("Evidences:")
			for _, ev := range task.Evidences {
				fmt.Println("  Comment :", ev.Comment)
				fmt.Println("  ID evidence :", ev.EvidenceID)
				fmt.Println(" FilePath   :", ev.FilePath)
				fmt.Println(" CreatedAt   :", ev.CreatedAt)
			}
		}
		fmt.Println("Assigned Projects:")

		for _, asp := range pr.AssignedProjects {
			fmt.Println("  ProjectID :", asp.ProjectID)
			fmt.Println("  TeamID :", asp.TeamID)
		}



	}

	}

	fmt.Println("PROBANDO USERS-------------------LISNTADO TODOS:")
	
	rp:=repository.NewGormUserRepository(db)


	userFind,err:=rp.GetUserByID(context.Background(),"USR0")
	if err!=nil{
		panic(err)
	}
	fmt.Println(userFind)
	/*
	err=rp.DeleteUser(context.Background(),"USR4")
	if err!=nil{
		panic(err)
	}*/
	

/*
	fmt.Println("------------------------PROJECTS------------------------")
	for _, pr := range projects {
		fmt.Println("PRoject ID:", pr.ProjectID)
		fmt.Println("Name:", pr.Name)
		fmt.Println("Description:", pr.Description)
		fmt.Println("USerID:", pr.CreatedBy)

	}*/

}
