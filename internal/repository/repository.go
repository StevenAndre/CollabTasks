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
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByUsername(ctx context.Context,username string) (*entity.User, error)
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
	DeleteProject(ctx context.Context, proj_id string) error
}

type TaskRepository interface{
	CreateTask(ctx context.Context,task *entity.Task)error
	GetAllTasksByProject(ctx context.Context,proj_id string) (*[]entity.Task, error)
	GetAllTasksByUser(ctx context.Context,user_id string) (*[]entity.Task, error)
	CompleteTask(ctx context.Context,task_id string) error
	GetTasksUncompleted(ctx context.Context,proj_id string) (*[]entity.Task, error)
	GetTasksCompleted(ctx context.Context,proj_id string) (*[]entity.Task, error)
}


type EvidenceRepository interface{
	SaveEvidence(ctx context.Context, evidence *entity.Evidence)error
	GetEvidenceById(ctx context.Context, evidence_id string)(*entity.Evidence,error)
	GetEvidencesByTask(ctx context.Context, task_id string)(*[]entity.Evidence,error)
	DeleteEvidence(ctx context.Context,evidence_id string)error
}

type PositionRepository interface{
	CreatePosition(ctx context.Context,position *entity.Position) error
	GetPositionById(ctx context.Context,position_id int16)(*entity.Position,error)
	GetAllPositions(ctx context.Context)(*[]entity.Position,error)
	UpdatePosition(ctx context.Context, position_id int16, newPosition *entity.Position) error
	DeletePosition(ctx context.Context, position_id int16) error
}

type TeamRepository interface{
	CreateTeam(ctx context.Context, team *entity.Team) error
	GetTeamById(ctx context.Context,team_id string)(*entity.Team,error)
	DeleteTeam(ctx context.Context,team_id string) error
}

type TeamMemberRepository interface{
	AssignTeam(ctx context.Context, team_member *entity.TeamMember)error
	RemoveFromTeam(ctx context.Context, member_id string, team_id string) error
	GetMembersByTeam(ctx context.Context,team_id string)(*[]entity.User,error)
	UpdatePositionTeam(ctx context.Context,update_position int16, id_team,id_member string) error
}


type AssignedProjectRepository interface{
	AssignProjectToTeam(ctx context.Context, assigned_project *entity.AssignedProject) error
	GetProjectsFromTeam(ctx context.Context,id_team string)(*[]entity.Project,error)
	RemoveProjectFromTeam(ctx context.Context,id_team,id_project string) error
}

type PrivateMessageRepsoitory interface{
	SaveMessage(ctx context.Context,message *entity.PrivateMessage)error
	GetMessagesBetweenUsers(ctx context.Context, sender_id ,receiver_id string)(*[]entity.PrivateMessage,error)
	DeleteMessage(ctx context.Context,pms_id string)error
	ReadMessage(ctx context.Context,pms_id string) error //update is_read from private messages
}

type TeamMessageRepository interface {
	SendMessageToTeam(ctx context.Context, message *entity.TeamMessage) error
	GetMessagesFromTeam(ctx context.Context,team_id string)(*[]entity.TeamMessage,error)
	DeleteMessage(ctx context.Context,tms_id string)error
}

