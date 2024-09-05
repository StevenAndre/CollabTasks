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

type NotificationService interface{
	SendNotification(ctx context.Context, id_user string, notification *dtos.NotificationDto ) error
	ListNotificationByUSer(ctx context.Context,id_user string)(*[]model.Notification,error)
	ReadNotification(ctx context.Context, id_notification string) error
	DeleteNotification(ctx context.Context, id_motification string)error
}

type TaskService interface{
	RegisterTask(ctx context.Context,taskDto dtos.TaskDto, userID,projectId string)error
	ListTasksByproject(ctx context.Context,projectId string)(*[]model.TaskU,error)
	ListTasksByUser(ctx context.Context,userID string)(*[]model.Task,error)
	CompleteTask(ctx context.Context, taskId string)error
	ListTasksUncompleted(ctx context.Context,projectId string) (*[]model.TaskU,error)
	ListTasksCompleted(ctx context.Context,projectId string) (*[]model.TaskU,error)
}

type EvidenceService interface {
	RegisterEvidence(ctx context.Context,evidenceDto dtos.NotificationDto,taskID string, file *multipart.FileHeader) error
	ShowEvidenceByID(ctx context.Context,evidenceID string)(*model.Evidence,error)
	ViewEvidencesByTask(ctx context.Context,taskID string)(*[]model.Evidence,error)
	DeleteEvidence(ctx context.Context,evidenceID string) error
}

type PositionsService interface{
	RegisterPosition(ctx context.Context, position dtos.PositionDto)error
	GetPositionByID(ctx  context.Context,positionID int16)(*model.Position,error)
	GetAllPositions(ctx context.Context) (*[]model.Position,error)
	UpdatePosition(ctx context.Context,positionID int16,position dtos.PositionDto)error
	RemovePosition(ctx context.Context,positionID int16)error
}

type TeamService interface{
	RegisterTeam(ctx context.Context,teamDto dtos.TeamDto)error
	GetTeamByID(ctx context.Context, teamID string) (*model.Team,error)
	DeleteTeam(ctx context.Context, teamID string)error
}

type TeamMemberService interface{
	AssignTeam(ctx context.Context,teammemberDto dtos.TeamMemberDto) error
	RemoveTeamMember(ctx context.Context,memberID, teamID string)error
	GetMemberByTeam(ctx context.Context, teamID string)(*[]model.User,error)
	UpdatePositionTeam(ctx context.Context,update_position int16, id_team,id_member string) error
}