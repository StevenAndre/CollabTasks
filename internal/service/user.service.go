package service

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/StevenAndre/collabtasks/encryption"
	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/mappers"
	"github.com/StevenAndre/collabtasks/internal/model"
	"github.com/StevenAndre/collabtasks/internal/repository"
	"github.com/StevenAndre/collabtasks/utils"
)

var mapper mappers.UserMapper
var (
	UsernameAlreadyExists = errors.New("The username already exists")
	EmailAlreadyExists    = errors.New("The email already exists")
)

type UserServiceImpl struct {
	repouser repository.UserRepository
}

func NewUserService(usrp repository.UserRepository) UserService {
	return &UserServiceImpl{
		repouser: usrp,
	}
}

func (us *UserServiceImpl) RegisterUser(ctx context.Context, userDto dtos.UserResisterDto, file *multipart.FileHeader) error {

	err := UsernameOrEmailAlreadyExists(ctx, userDto.Username, userDto.Email, us.repouser)
	if err != nil {
		return err
	}
	id := utils.GeneradorIds(4)
	var profilepath string
	psswenb, err := encryption.Encrypt([]byte(userDto.Password))
	if err != nil {
		return err
	}
	passwto64 := encryption.ToBase64(psswenb)
	if file == nil {
		profilepath = "defaul-profile.jpg"
	} else {
		profilepath, err = saveFile(file)
		if err != nil {
			return err
		}
	}
	u := entity.User{
		UserID:       id,
		Username:     userDto.Username,
		Name:         userDto.Name,
		Lastname:     userDto.Lastname,
		Email:        userDto.Email,
		Password:     passwto64,
		ProfilePhoto: profilepath,
	}
	err = us.repouser.SaveUser(ctx, &u)
	return err
}

func (us *UserServiceImpl) GetAllUsers(ctx context.Context) (*[]model.User, error) {
	users := &[]model.User{}
	usesent, err := us.repouser.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range *usesent {
		*users = append(*users, *mapper.FromEntityToModel(&u))
	}
	return users, nil
}

func (us *UserServiceImpl) GetUserById(ctx context.Context, userID string) (*model.User, error) {
	user, err := us.repouser.GetUserByID(ctx, userID)
	
	if err != nil {
		return nil, err
	}
	return mapper.FromEntityToModel(user), nil
}

func (us *UserServiceImpl) UpdateUser(ctx context.Context, userDto dtos.UserUpdateDto, userID string, file *multipart.FileHeader) error {
	usDB, err := us.repouser.GetUserByID(ctx, userID)
	
	if err != nil {
		return err
	}
	
	
	
	if usDB.Email != userDto.Email{
		usex,err:=us.repouser.GetUserByEmail(ctx,userDto.Email)
		if err!=nil{
			return err
		}
		if usex!=nil{
			return EmailAlreadyExists
		}
	}

	if usDB.Username != userDto.Username{
		usex,err:=us.repouser.GetUserByUsername(ctx,userDto.Username)
		if err!=nil{
			return err
		}
		if usex!=nil{
			return UsernameAlreadyExists
		}
	}

	psswenb, err := encryption.Encrypt([]byte(userDto.Password))
	if err != nil {
		return err
	}
	passwto64 := encryption.ToBase64(psswenb)

	 profilepath := usDB.ProfilePhoto
	
	
		if file != nil {
			err=deleteFile(profilepath)
			if err != nil {
				return err
			}
			profilepath, err = saveFile(file)
			if err != nil {
				return err
			}
		}
			u := entity.User{
				UserID:       usDB.UserID,
				Username:     userDto.Username,
				Name:         userDto.Name,
				Lastname:     userDto.Lastname,
				Email:        userDto.Email,
				Password:     passwto64,
				ProfilePhoto: profilepath,
			}
			err = us.repouser.UpdateUser(ctx, &u)
		
	return err 
}

func (us *UserServiceImpl) DeleteUser(ctx context.Context, userID string) error {
	usDB, err := us.repouser.GetUserByID(ctx, userID)
	if err!=nil{
		return err
	}
	err=deleteFile(usDB.ProfilePhoto)
	if err!=nil{
		return err
	}
	return us.repouser.DeleteUser(ctx,userID)
}



func UsernameOrEmailAlreadyExists(ctx context.Context, username, email string, rp repository.UserRepository) error {
	usEx, err := rp.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	if usEx != nil {
		return EmailAlreadyExists
	}
	usEx, err = rp.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}
	if usEx != nil {
		return UsernameAlreadyExists
	}
	return nil
}
