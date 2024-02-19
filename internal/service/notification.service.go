package service

import (
	"context"
	"errors"
	"time"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/mappers"
	"github.com/StevenAndre/collabtasks/internal/model"
	"github.com/StevenAndre/collabtasks/internal/repository"
	"github.com/StevenAndre/collabtasks/utils"
)

var notMapper mappers.NotificationMapper

type NotificationServiceImpl struct {
	nr repository.NotificationRepository
	ur repository.UserRepository
}

func NewNotificationService(nr *repository.NotificationRepository, ur *repository.UserRepository) NotificationService {
	return &NotificationServiceImpl{
		nr: *nr,
		ur: *ur,
	}
}

func (ns *NotificationServiceImpl) SendNotification(ctx context.Context, id_user string, notification *dtos.NotificationDto) error {
	us, err := ns.ur.GetUserByID(ctx, id_user)
	if err != nil {
		return err
	}
	if us == nil {
		return errors.New("User not found")
	}
	id := utils.GeneradorIds(10)
	ntf := &entity.Notification{
		NotificationID: id,
		Issue:          notification.Issue,
		Issued:         time.Now(),
		Message:        notification.Message,
		Isread:         false,
		UserID:         us.UserID,
	}
	err = ns.nr.CreateNotification(ctx, ntf)
	return err

}
func (ns *NotificationServiceImpl) ListNotificationByUSer(ctx context.Context, id_user string) (*[]model.Notification, error) {

	notifications := &[]model.Notification{}

	notifDB, err := ns.nr.GetAllNotifiationsByUser(ctx, id_user)
	if err != nil {
		return nil, err
	}
	for _, v := range *notifDB {
		*notifications = append(*notifications, *notMapper.FromEntityToModel(&v))
	}

	return notifications, nil
}
func (ns *NotificationServiceImpl) ReadNotification(ctx context.Context, id_notification string) error {
	err:=ns.nr.MarkAsRead(ctx,id_notification)
	return err
}
func (ns *NotificationServiceImpl) DeleteNotification(ctx context.Context, id_motification string) error {
	err:=ns.nr.DeleteNotification(ctx,id_motification)
	return err
}
