package repository

import (
	"context"
	"errors"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormProjectRepository struct {
	db *gorm.DB
}

func NewGormProjectRepository(db *gorm.DB) ProjectRepository {
	return &GormProjectRepository{
		db: db,
	}
}

func (pr *GormProjectRepository) AddProject(ctx context.Context, project *entity.Project) error {
	result := pr.db.Create(project)
	return result.Error
}

func (pr *GormProjectRepository) GetProjectsByUser(ctx context.Context, user_id string) (*[]entity.Project, error) {

	projects := &[]entity.Project{}

	if user_id == "" {
		return nil, errors.New("ID bad format")
	}
	result := pr.db.Where("create_by = ?", user_id).Find(&projects)

	if result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}

func (pr *GormProjectRepository) GetProjectByID(ctx context.Context, proj_id string) (*entity.Project, error) {

	if proj_id == "" {
		return nil, errors.New("ID bad format")
	}
	project := &entity.Project{}
	result := pr.db.Where("id_project = ?", proj_id).First(project)

	if result.Error != nil {
		return nil, result.Error
	}

	return project, nil
}

func (pr *GormProjectRepository) DeleteProject(ctx context.Context, proj_id string) error {
	project := entity.Project{
		ProjectID: proj_id,
	}
	resutl := pr.db.Delete(&project)
	return resutl.Error
}
