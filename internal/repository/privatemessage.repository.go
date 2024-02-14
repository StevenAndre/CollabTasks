package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormPrivateMessageRepository struct {
	db *gorm.DB
}

func NewGormPrivateMessageRepository(db *gorm.DB)PrivateMessageRepsoitory{
	return &GormPrivateMessageRepository{db: db}
}

func (pmr *GormPrivateMessageRepository) SaveMessage(ctx context.Context,message *entity.PrivateMessage)error{
	result:=pmr.db.Create(message)
	return result.Error
}

func (pmr *GormPrivateMessageRepository) GetMessagesBetweenUsers(ctx context.Context, sender_id ,receiver_id string)(*[]entity.PrivateMessage,error){
	messages:=&[]entity.PrivateMessage{}
	result:=pmr.db.Where("sender_id = ?",sender_id).
	Where("receiver_id = ?",receiver_id).
	Order("sent_date DESC").
	Find(messages)
	if result.Error != nil {
		return nil,result.Error
	}
	return messages,nil

}

func (pmr *GormPrivateMessageRepository) DeleteMessage(ctx context.Context,pms_id string)error{
	message:=&entity.PrivateMessage{
		PriMsgID: pms_id,
	}
	result := pmr.db.Delete(message)
	return result.Error
}

func (pmr *GormPrivateMessageRepository) ReadMessage(ctx context.Context,pms_id string) error{
	message:=&entity.PrivateMessage{
		PriMsgID: pms_id,
	}
	result:=pmr.db.First(message)
	if result.Error!=nil {
		return result.Error
	}
	message.IsRead=true
	result=pmr.db.Save(message)
	return result.Error
}