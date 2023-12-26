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
}