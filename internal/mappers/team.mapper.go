package mappers

import (
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/model"
)

func (tm *TeamMapper) FromEntityToModel(team *entity.Team) *model.Team{
	return  &model.Team {
		TeamID: team.TeamID,
		Name:    team.Name,
	}
}