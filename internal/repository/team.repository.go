package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormTeamRepository struct {
	db *gorm.DB
}

func NewGormTeamRepository(db *gorm.DB) TeamRepository{
	return &GormTeamRepository{
		db: db,
	}
}

func (tr *GormTeamRepository) CreateTeam(ctx context.Context, team *entity.Team) error{
	result:=tr.db.Create(team)
	return result.Error
}

func (tr *GormTeamRepository) GetTeamById(ctx context.Context,team_id string)(*entity.Team,error){
	team:=&entity.Team{
		TeamID: team_id,
	}
	result:=tr.db.First(team)
	if result.Error!=nil{
		return nil, result.Error
	}
	
	return team,nil
}

func (tr *GormTeamRepository) DeleteTeam(ctx context.Context,team_id string) error{
	return tr.db.Delete(&entity.Team{
		TeamID: team_id,
		}).Error
}