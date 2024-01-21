package api

import (
	"github.com/StevenAndre/collabtasks/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UserController struct {
	sv service.UserService
	datavalid *validator.Validate
}

var dataValis =validator.New()

func NewUserController(e *echo.Echo, sv *service.UserService)*UserController{
	userController:=&UserController{
		sv: *sv,
		datavalid: dataValis,
	}
	userController.RegisterRoutes(e)
	return userController
}

type ProjectController struct{
	sv service.ProjectService
	datavalid *validator.Validate
}

func NewProjectController(e *echo.Echo, sv *service.ProjectService)*ProjectController{
	ProjectController:= &ProjectController{
		sv: *sv,
		datavalid: dataValis,
	}
	ProjectController.RegisterRoutes(e)
	return ProjectController
}



func Start(e *echo.Echo,address string)error{
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType},
		AllowCredentials: true,
	}))
	return e.Start(address)
}
