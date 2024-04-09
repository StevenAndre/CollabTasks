package service

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/mappers"
	"github.com/StevenAndre/collabtasks/internal/model"
	"github.com/StevenAndre/collabtasks/internal/repository"
	"github.com/StevenAndre/collabtasks/utils"
)

var teamMapper mappers.TeamMapper

type TeamServiceImp struct {
	teamRepo repository.TeamRepository
}

func NewTeamService(teamRp *repository.TeamRepository)TeamService{
	return &TeamServiceImp{
		teamRepo: *teamRp,
	}
}


func (t * TeamServiceImp)RegisterTeam(ctx context.Context,teamDto dtos.TeamDto)error{
	team:=&entity.Team{
		TeamID: utils.GeneradorIds(8),
		Name: teamDto.Name,
	}
	err:=t.teamRepo.CreateTeam(ctx,team)
	return err
}
func (t * TeamServiceImp)GetTeamByID(ctx context.Context, teamID string) (*model.Team,error){
	team,err:=t.teamRepo.GetTeamById(ctx,teamID)
	if err!=nil{
		return nil,err
	}
	return teamMapper.FromEntityToModel(team),nil
}
func (t * TeamServiceImp)DeleteTeam(ctx context.Context, teamID string)error{
	err:=t.teamRepo.DeleteTeam(ctx,teamID)
	return err
}