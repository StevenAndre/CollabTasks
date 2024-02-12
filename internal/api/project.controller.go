package api

import (
	"net/http"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/labstack/echo/v4"
)

func (pc *ProjectController) CreateProject(c echo.Context) error {

	ctx := c.Request().Context()
	data := dtos.ProjectDto{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = pc.datavalid.Struct(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = pc.sv.AddProject(ctx, data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	return c.NoContent(http.StatusOK)

}

func (pc *ProjectController) ListProjectsOfUser(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Param("id_user")
	if userId==""{
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "ID Bad format"})
	}

	projects, err := pc.sv.GetProjectsByUser(ctx, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	if projects == nil {
		return c.JSON(http.StatusOK, []any{})
	}

	return c.JSON(http.StatusOK, projects)
}

func (pc *ProjectController) ViewProjectByID(c echo.Context) error {
	ctx := c.Request().Context()
	projectId := c.Param("id_project")
	if projectId==""{
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "ID Bad format"})
	}
	project, err := pc.sv.GetProjectByID(ctx, projectId)
	if err != nil {
		return c.JSON(http.StatusNotFound, responseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, project)
}

func (pc *ProjectController) DeleteProject(c echo.Context) error {
	ctx := c.Request().Context()
	projectId := c.Param("id_project")
	err := pc.sv.DeleteProject(ctx, projectId)
	if err != nil {
		return c.JSON(http.StatusNotFound, responseMessage{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}
