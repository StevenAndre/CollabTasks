package mappers

import (
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/model"
)

func (pm * ProjectMapper)FromEntityToModel(project *entity.Project)*model.Project{
	return &model.Project {
		ProjectID: project.ProjectID,
		Name:      project.Name,
		Description: project.Description,
		CreatedAt: project.CreatedAt,
		CreatedBy: project.CreatedBy,
	}
}