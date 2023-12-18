package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository{
	return &GormUserRepository{
		db:db,
	}
}


func (ur *GormUserRepository) SaveUser(ctx context.Context, user *entity.User) error{
	result:=ur.db.Create(user)
	return result.Error
}
func (ur *GormUserRepository) UpdateUser(ctx context.Context, user *entity.User) error{
	result:=ur.db.Save(user)
	return result.Error
}
func (ur *GormUserRepository) GetUserByID(ctx context.Context, id_user string) (*entity.User, error){
	user:=&entity.User{
		UserID: id_user,
	}
	result:=ur.db.First(user)
	
	return user,result.Error
}
func (ur *GormUserRepository) GetAllUsers(ctx context.Context)(*[]entity.User,error){
	users := []entity.User{}
	result:= ur.db.Find(&users)
	return &users,result.Error
}
func (ur *GormUserRepository) DeleteUser(ctx context.Context, id_user string) error{
	return nil
}