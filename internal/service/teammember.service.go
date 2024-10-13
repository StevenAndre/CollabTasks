package service

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/model"
	"github.com/StevenAndre/collabtasks/internal/repository"
)

type TeamMemberServiceImpl struct {
	teamMemRepo repository.TeamMemberRepository
}




func NewTeamMemberService(teamMemRepo *repository.TeamMemberRepository)TeamMemberService{
	return &TeamMemberServiceImpl{
		teamMemRepo: *teamMemRepo,
	}
}

func (tm *TeamMemberServiceImpl)AssignTeam(ctx context.Context,teammemberDto dtos.TeamMemberDto) error{
	
	newTeamMember:=entity.TeamMember{
		 UserID: teammemberDto.UserID,
		 TeamID: teammemberDto.TeamID,
		 PositionID: teammemberDto.PositionID,
	}
	 err:=tm.teamMemRepo.AssignTeam(ctx,&newTeamMember)
	return err
}
func (tm *TeamMemberServiceImpl)RemoveTeamMember(ctx context.Context,memberID, teamID string)error{
	return tm.teamMemRepo.RemoveFromTeam(ctx,memberID,teamID)
}
func (tm *TeamMemberServiceImpl)GetMemberByTeam(ctx context.Context, teamID string)(*[]model.User,error){


	usersEnty,err:=tm.teamMemRepo.GetMembersByTeam(ctx,teamID)
	if err!=nil{
		return nil,err
	}

	users:=&[]model.User{}
	for _,ue := range *usersEnty{
		*users=append(*users, *mapper.FromEntityToModel(&ue))
	}

	return users,nil
}
func (tm *TeamMemberServiceImpl)UpdatePositionTeam(ctx context.Context,update_position int16, id_team,id_member string) error{
	return tm.teamMemRepo.UpdatePositionTeam(ctx,update_position,id_team,id_member)
}
