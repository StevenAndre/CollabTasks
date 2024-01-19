package service

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/StevenAndre/collabtasks/settings"
	"github.com/google/uuid"
)


var pathImages string=settings.GetPathImages()



func saveFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {

		return "", err
	}
	defer file.Close()

	filename := fmt.Sprintf("%s-%s.jpg", uuid.NewString(), fileHeader.Filename)
	dst, err := os.Create(pathImages + filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}
	return filename, nil

}

func deleteFile(filename string) error {
	if filename == "perfil-defecto.png" {
		return nil
	}
	err := os.Remove(pathImages + filename)
	if err != nil {
		_, s := err.(*os.PathError)
		if s {
			log.Println(err.Error())
			return nil
		} else {
			return err
		}

	}
	return nil
}
