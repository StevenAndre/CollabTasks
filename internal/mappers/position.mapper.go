package mappers

import (
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/model"
)

func (pm *PositionMapper) FromEntityToModel(position *entity.Position)*model.Position{
	return &model.Position{
		PositionID: position.PositionID,
		Name: position.Name,
		Description: position.Description,
	}
}