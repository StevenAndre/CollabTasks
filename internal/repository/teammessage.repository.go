package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormTeamMessageRepository struct {
	db *gorm.DB
}

func NewGormTeamMessageRepository(db *gorm.DB)TeamMessageRepository{
	return &GormTeamMessageRepository{db: db}
}


func  (r *GormTeamMessageRepository) SendMessageToTeam(ctx context.Context, message *entity.TeamMessage) error{
	result:=r.db.Create(message)
	return result.Error
}

func  (r *GormTeamMessageRepository) GetMessagesFromTeam(ctx context.Context,team_id string)(*[]entity.TeamMessage,error){
	messages:=&[]entity.TeamMessage{}
	result:=r.db.Where("team_id = ?",team_id).
	Order("send_date DESC").
	Find(messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages,  nil
}

func  (r *GormTeamMessageRepository) DeleteMessage(ctx context.Context,tms_id string)error{
	message:=&entity.TeamMessage{
		TeamMsgID: tms_id,
	}
	result:=r.db.Delete(message)
	return result.Error
}