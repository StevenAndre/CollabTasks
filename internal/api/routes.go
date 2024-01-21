package api

import "github.com/labstack/echo/v4"

func (uc *UserController) RegisterRoutes(e *echo.Echo){
	users:=e.Group("/users")
	users.POST("/create",uc.RegisterUser)
	users.GET("/all",uc.ListAllUsers)
	users.GET("/:id_user",uc.GetUserByID)
	users.PUT("/update/:id_user",uc.Updateuser)
	users.DELETE("/delete/:id_user", uc.DeleteUser)
}


func (pc *ProjectController)RegisterRoutes(e *echo.Echo){
	projects:=e.Group("/projects")
	projects.POST("/create",pc.CreateProject)
	projects.GET("/user/:id_user",pc.ListProjectsOfUser)
	projects.GET("/:id_project",pc.ViewProjectByID)
	projects.DELETE("/delete/:id_project",pc.DeleteProject)
	
}
