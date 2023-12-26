package api

import "github.com/labstack/echo/v4"

func (uc *UserController) RegisterRoutes(e *echo.Echo){
	users:=e.Group("/users")
	users.POST("/create",uc.RegisterUser)
	users.GET("/all",uc.ListAllUsers)
}