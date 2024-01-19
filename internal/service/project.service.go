package service

import (
	"context"
	"time"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/mappers"
	"github.com/StevenAndre/collabtasks/internal/model"
	"github.com/StevenAndre/collabtasks/internal/repository"
	"github.com/StevenAndre/collabtasks/utils"
)

var pj_mapper mappers.ProjectMapper

type ProjectServiceImpl struct {
	rppj repository.ProjectRepository
	rpus repository.UserRepository
}

func NewProjectService(rp *repository.ProjectRepository, rpu *repository.UserRepository) ProjectService {

	return &ProjectServiceImpl{
		rppj: *rp,
		rpus: *rpu,
	}
}

func (ps *ProjectServiceImpl) AddProject(ctx context.Context, project_dto dtos.ProjectDto) error {
	us, err := ps.rpus.GetUserByID(ctx, project_dto.CreatedBy)
	if err != nil {
		return err
	}

	id := utils.GeneradorIds(8)
	project := entity.Project{
		ProjectID:   id,
		Name:        project_dto.Name,
		CreatedAt:   time.Now(),
		CreatedBy:   us.UserID,
		Description: project_dto.Description,
	}

	return ps.rppj.AddProject(ctx, &project)
}
func (ps *ProjectServiceImpl) GetProjectsByUser(ctx context.Context, user_id string) (*[]model.Project, error) {
	projects := &[]model.Project{}

	projectsBD, err := ps.rppj.GetProjectsByUser(ctx, user_id)
	if err != nil {
		return nil, nil
	}
	for _, p := range *projectsBD {
		*projects = append(*projects, *pj_mapper.FromEntityToModel(&p))
	}

	return projects, nil
}
func (ps *ProjectServiceImpl) GetProjectByID(ctx context.Context, proj_id string) (*model.Project, error) {
	project:=&model.Project{}

	pjDB,err:= ps.rppj.GetProjectByID(ctx,proj_id)
	if err!=nil{
		return nil, err
	}

	project=pj_mapper.FromEntityToModel(pjDB)
	return project, nil
}


func (ps *ProjectServiceImpl) DeleteProject(ctx context.Context, proj_id string) error {
	err:=ps.rppj.DeleteProject(ctx,proj_id)
	return err
}
