package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormNotificationRepository struct {
	db *gorm.DB
}

func NewGormNotificationRepository(db *gorm.DB) NotificationRepository{
	return &GormNotificationRepository{
		db: db,
	}
}


func (nr * GormNotificationRepository)CreateNotification(ctx context.Context,notification *entity.Notification)error{
	result:=nr.db.Create(notification)
	return result.Error
}
func (nr * GormNotificationRepository) GetAllNotifiationsByUser(ctx context.Context,user_id string)(*[]entity.Notification,error){
	notifications:=&[]entity.Notification{}
	result:=nr.db.Where("id_user = ?",user_id).Find(notifications)
	if result.Error!=nil{
		return  nil,result.Error
	}
	return notifications,nil
}
func (nr * GormNotificationRepository) MarkAsRead(ctx context.Context,notif_id string)error{
	result:=nr.db.Model(&entity.Notification{}).Where("id_notification = ?",notif_id).Update("is_read",true)
	return result.Error
}
func (nr * GormNotificationRepository) DeleteNotification(ctx context.Context,notif_id string)error{
	notification:=entity.Notification{
		NotificationID: notif_id,
	}
	result:=nr.db.Delete(&notification)
	return result.Error
}
