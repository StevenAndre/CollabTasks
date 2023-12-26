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

func NewUserController(e *echo.Echo, sv service.UserService)*UserController{
	userController:=&UserController{
		sv: sv,
		datavalid: validator.New(),
	}
	userController.RegisterRoutes(e)
	return userController
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
