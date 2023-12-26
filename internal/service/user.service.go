package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/StevenAndre/collabtasks/encryption"
	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/model"
	"github.com/StevenAndre/collabtasks/internal/repository"
	"github.com/StevenAndre/collabtasks/utils"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	repouser repository.UserRepository
}

func NewUserService(usrp repository.UserRepository) UserService {
	return &UserServiceImpl{
		repouser: usrp,
	}
}


func(us *UserServiceImpl)RegisterUser(ctx context.Context,userDto dtos.UserResisterDto,file *multipart.FileHeader) error{
	id:=utils.GeneradorIds(4)
	var profilepath string
	psswenb,err:=encryption.Encrypt([]byte(userDto.Password))
	if err!=nil{
		return err
	}
	passwto64:=encryption.ToBase64(psswenb)
	if file==nil{
		profilepath="defaul-profile.jpg"
	}else{
		profilepath,err=saveFile(file)
		if err!=nil{
			return err
		}
	}
	u:=entity.User{
		UserID: id,
		Username: userDto.Username,
		Name:   userDto.Name,
		Lastname: userDto.Lastname,
		Email: userDto.Email,
		Password: passwto64,
		ProfilePhoto: profilepath,
	}
	err=us.repouser.SaveUser(ctx,&u)
	return err
}

func (us *UserServiceImpl)GetAllUsers(ctx context.Context) (*[]model.User,error){
	users:=&[]model.User{}
	usesent,err:=us.repouser.GetAllUsers(ctx)
	if err!= nil{
		return nil,err
	}

	for _,u := range *usesent{
		*users=append(*users,model.User{
			UserID: u.UserID,
			Name: u.Name,
			Lastname: u.Lastname,
			Username: u.Username,
			Email: u.Email,
			ProfilePhoto: u.ProfilePhoto,
		})
	}
	return users,nil
}




const pathImages="C:/Users/steve/OneDrive/Escritorio/Proyectos programcion/GoProyectos/CollabTasks/images/"

func saveFile(fileHeader *multipart.FileHeader) (string,error) {
	file,err:=fileHeader.Open()
	if err!=nil{
	
		return "",err
	}
	defer file.Close()

	filename:=fmt.Sprintf("%s-%s.jpg",uuid.NewString(),fileHeader.Filename)
	dst,err:=os.Create(pathImages+filename)
	if err!=nil{
		return "",err
	}
	defer dst.Close()
	_,err=io.Copy(dst,file)
	if err != nil {
		return "",err
	}
	return filename,nil


}