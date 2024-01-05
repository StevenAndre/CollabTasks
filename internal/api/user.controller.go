package api

import (
	"encoding/json"
	"net/http"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (uc *UserController) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	data := dtos.UserResisterDto{}
	userJSON := c.FormValue("user")
	

	if err := json.Unmarshal([]byte(userJSON), &data); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err := uc.datavalid.Struct(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	fileHeader, err := c.FormFile("photo")

	if err != nil {
		err = uc.sv.RegisterUser(ctx, data, nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
		}
	} else {
		err = uc.sv.RegisterUser(ctx, data, fileHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
		}
	}
	return c.JSON(http.StatusCreated, nil)

}


func  (uc *UserController) Updateuser(c echo.Context) error {
	ctx := c.Request().Context()
	data := dtos.UserUpdateDto{}
	userJSON := c.FormValue("user")
	idUser:=c.Param("id_user")

	if err := json.Unmarshal([]byte(userJSON), &data); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err := uc.datavalid.Struct(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	fileHeader, err := c.FormFile("photo")

	if err != nil {
		err = uc.sv.UpdateUser(ctx, data,idUser, nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
		}
	} else {
		err = uc.sv.UpdateUser(ctx, data,idUser, fileHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
		}
	}
	return c.JSON(http.StatusOK, nil)
}



func (uc *UserController)GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Param("id_user")
	us,err:=uc.sv.GetUserById(ctx,userId)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,responseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK,*us)

}

func (uc *UserController)DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Param("id_user")
	err:=uc.sv.DeleteUser(ctx,userId)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,responseMessage{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}




func (uc *UserController) ListAllUsers(c echo.Context) error {
	users, err := uc.sv.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, responseMessage{"No users found."})
	}
	return c.JSON(http.StatusOK, *users)
}
