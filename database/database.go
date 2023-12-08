package database

import (
	"context"
	"fmt"

	"github.com/StevenAndre/collabtasks/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func NewDatabase(ctx context.Context,s *settings.Settings)(* gorm.DB, error) {
	var datasource = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.DB.Host, s.DB.Port, s.DB.User, s.DB.Password, s.DB.Name)

		DB, err := gorm.Open(postgres.Open(datasource), &gorm.Config{
			
		})
		if err!=nil{
			return nil,err
		}
		return DB,err
}



func CloseDB() error {
    db, err := DB.DB() // devuelve el pool de conexiones 
    if err != nil {
        return err
    }
    
    return db.Close()
}

type Namer interface {
	TableName(table string) string
}