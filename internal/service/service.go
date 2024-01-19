package service

import (
	"context"
	"mime/multipart"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/StevenAndre/collabtasks/internal/model"
)

type UserService interface {
	RegisterUser(ctx context.Context,userDto dtos.UserResisterDto,file *multipart.FileHeader) error
	GetAllUsers(ctx context.Context) (*[]model.User,error)
	GetUserById(ctx context.Context,userID string)(*model.User,error)
	UpdateUser(ctx context.Context,userDto dtos.UserUpdateDto, userID string,file *multipart.FileHeader)error
	DeleteUser(ctx context.Context,userID string)error
}

type ProjectService interface{
	AddProject(ctx context.Context, project_dto dtos.ProjectDto)error
	GetProjectsByUser(ctx context.Context,user_id string) (*[]model.Project ,error)
	GetProjectByID(ctx context.Context,proj_id string) (*model.Project, error)
	DeleteProject(ctx context.Context, proj_id string) error
}