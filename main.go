package main

import (
	"context"
	

	"github.com/StevenAndre/collabtasks/database"
	"github.com/StevenAndre/collabtasks/internal/api"
	"github.com/StevenAndre/collabtasks/internal/repository"
	"github.com/StevenAndre/collabtasks/internal/service"
	"github.com/StevenAndre/collabtasks/settings"
	"github.com/labstack/echo/v4"
)

func main() {

	s, err := settings.NewSettings()
	if err != nil {
		panic(err)
	}

	db, err := database.NewDatabase(context.Background(), s)
	if err != nil {
		panic(err)
	}
	userRepo:=repository.NewGormUserRepository(db)
	userService:=service.NewUserService(&userRepo)
	projectRepo:=repository.NewGormProjectRepository(db)
	projectService:=service.NewProjectService(&projectRepo,&userRepo)
	e:=echo.New()
	_=api.NewUserController(e,&userService)
	_=api.NewProjectController(e,&projectService)
	api.Start(e,":8080")

}
