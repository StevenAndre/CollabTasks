package mappers

import (
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/model"
)

func (nm *NotificationMapper) FromEntityToModel(notification *entity.Notification) *model.Notification{
	return &model.Notification{
		NotificationID: notification.NotificationID,
		Issue: notification.Issue,
		Message: notification.Message,
		Issued: notification.Issued,
		Isread: notification.Isread,
		UserID: notification.UserID,
	}
}