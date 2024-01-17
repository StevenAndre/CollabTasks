package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormProjectRepository struct {
	db *gorm.DB
}


func NewGormProjectRepository(db *gorm.DB) ProjectRepository {
	return  &GormProjectRepository{
		db: db,
	}
}


func (pr *GormProjectRepository) AddProject(ctx context.Context,project *entity.Project) error{
	return nil
}
func (pr *GormProjectRepository) GetProjectsByUser(ctx context.Context,user_id string) (*[]entity.Project ,error){
	return nil, nil
}
func (pr *GormProjectRepository) GetProjectByID(ctx context.Context,proj_id string) (*entity.Project, error){
	return nil , nil
}
func (pr *GormProjectRepository) DeleteProject(ctx context.Context, proj_id string) error{
	return nil
}