package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id_user string) (*entity.User, error)
	GetAllUsers(ctx context.Context)(*[]entity.User,error)
	DeleteUser(ctx context.Context, id_user string) error
}

type NotificationRepository interface{

	CreateNotification(ctx context.Context,notification *entity.Notification)error
	GetAllNotifiationsByUser(ctx context.Context,user_id string)(*[]entity.Notification,error)
	


}