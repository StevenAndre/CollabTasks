package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormAssignedProjectRepository struct {
	db *gorm.DB
}

func NewGormAssignedProjectRepository(db *gorm.DB) AssignedProjectRepository{
	return  &GormAssignedProjectRepository{db: db}
}

func (ap *GormAssignedProjectRepository) AssignProjectToTeam(ctx context.Context, assigned_project *entity.AssignedProject) error{
	result:=ap.db.Create(assigned_project)
	return result.Error
}
func (ap *GormAssignedProjectRepository) GetProjectsFromTeam(ctx context.Context,id_team string)(*[]entity.Project,error){
	projects:=&[]entity.Project{}
	result:=ap.db.Table("projects").
	Joins("INNER JOIN assigned_project ON projects.id_project=assigned_project.id_project").
	Where("assigned_project.id_team = ?", id_team).
	Find(projects)
	if result.Error != nil {
		return nil,result.Error
	}
	return projects,nil
}
func (ap *GormAssignedProjectRepository) RemoveProjectFromTeam(ctx context.Context,id_team,id_project string)error{
	assigProject:=&entity.AssignedProject{
		ProjectID: id_project,
		TeamID:  id_team,
	}
	result:=ap.db.Delete(assigProject)
	return result.Error
}