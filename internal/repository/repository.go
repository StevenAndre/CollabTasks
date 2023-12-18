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
	MarkAsRead(ctx context.Context,notif_id string)error
	DeleteNotification(ctx context.Context,notif_id string)error
}

type ProjectRepository interface{
	AddProject(ctx context.Context,project *entity.Project) error
	GetProjectsByUser(ctx context.Context,user_id string) (*[]entity.Project ,error)
	GetProjectByID(ctx context.Context,proj_id string) (*entity.Project, error)
}

type TaskRepository interface{
	CreateTask(ctx context.Context,task *entity.Task)error
	GetAllTasksByProject(ctx context.Context,proj_id string) (*[]entity.Task, error)
	GetAllTasksByUser(ctx context.Context,user_id string) (*[]entity.Task, error)
	CompleteTask(ctx context.Context,task_id string) error
	GetTasksUncompleted(ctx context.Context,proj_id string) (*[]entity.Task, error)
	GetTasksCompleted(ctx context.Context,proj_id string) (*[]entity.Task, error)
}
