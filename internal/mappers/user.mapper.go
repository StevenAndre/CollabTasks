package mappers

import (
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/model"
)

func (um *UserMapper) FromEntityToModel(user *entity.User)*model.User {
	return &model.User {
		UserID: user.UserID,
		Name: user.Name,
		Username: user.Username,
		Lastname: user.Lastname,
		Email: user.Email,
		ProfilePhoto: user.ProfilePhoto,
	}
}