package main

import (
	"context"
	"fmt"

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
	fmt.Println(s)

	db, err := database.NewDatabase(context.Background(), s)
	if err != nil {
		panic(err)
	}
	userRepo:=repository.NewGormUserRepository(db)
	userService:=service.NewUserService(&userRepo)
	e:=echo.New()
	_=api.NewUserController(e,userService)
	api.Start(e,":8080")

}
