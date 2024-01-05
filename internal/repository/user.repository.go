package repository

import (
	"context"
	"errors"
	"fmt"
	

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

var(
	UserNotFound=errors.New("")
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
	
	if id_user==""{
		return nil,errors.New("ID Bad Format")
	}
	user:=&entity.User{
		UserID: id_user,
	}
	result:=ur.db.Find(user)
	if result.RowsAffected==0{
		return nil,errors.New(fmt.Sprintf("User with ID: %s not found",id_user))
	}
	
	return user,result.Error
}
func (ur *GormUserRepository) GetAllUsers(ctx context.Context)(*[]entity.User,error){
	users := []entity.User{}
	result:= ur.db.Find(&users)
	return &users,result.Error
}
func (ur *GormUserRepository) DeleteUser(ctx context.Context, id_user string) error{
	rs:=ur.db.Delete(&entity.User{
		UserID: id_user,
	})
	if rs.RowsAffected==0{
		return nil
	}
	return rs.Error
}

func (ur *GormUserRepository)GetUserByEmail(ctx context.Context, email string) (*entity.User, error){
	user:=&entity.User{}
	rs:=ur.db.Where("email=?",email).Find(user)
	if rs.Error!=nil{
		return nil,rs.Error
	}
	if rs.RowsAffected==0{
		return nil,nil
	}
	return user,nil
}

func (ur *GormUserRepository)GetUserByUsername(ctx context.Context, username string) (*entity.User, error){
	user:=&entity.User{}
	rs:=ur.db.Where("username=? ",username,).Find(user)
	if rs.Error!=nil{
		return nil,rs.Error
	}
	if rs.RowsAffected==0{
		return nil,nil
	}
	return user,nil
}