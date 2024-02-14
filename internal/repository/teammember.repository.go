package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormTeamMemberRepository struct {
	db *gorm.DB
}

func NewGormTeamMemberRepository(db *gorm.DB) TeamMemberRepository{
	return &GormTeamMemberRepository{
		db: db,
	}
}

func(tm * GormTeamMemberRepository) AssignTeam(ctx context.Context, team_member *entity.TeamMember)error{
	result:=tm.db.Create(team_member)
	return result.Error
}

func(tm * GormTeamMemberRepository) RemoveFromTeam(ctx context.Context, member_id string, team_id string) error{
	teamMember:=&entity.TeamMember{
		UserID: member_id,
		TeamID: team_id,
	}
	result:=tm.db.Delete(teamMember)
	return result.Error
}

func(tm * GormTeamMemberRepository) GetMembersByTeam(ctx context.Context,team_id string)(*[]entity.User,error){
	members:=&[]entity.User{}
	result:=tm.db.Table("users").
	Select("users.*").
	Joins("JOIN team_members ON users.id_user = team_members.id_member").
	Where("team_members.id_team = ?", team_id).
	Find(&members)
	if  result.Error != nil{
		return  nil, result.Error
	}
	return members,nil
}

func(tm * GormTeamMemberRepository) UpdatePositionTeam(ctx context.Context,position_updated int16, id_team,id_member string)error{
	teamMember:=&entity.TeamMember{
		UserID: id_member,
		TeamID: id_team,
	}
	result := tm.db.First(teamMember)
	if  result.Error !=nil{
		return result.Error
	}

	teamMember.PositionID=position_updated
	result=tm.db.Save(teamMember)
	return result.Error
}