package service

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
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
	return nil
}
func (tm *TeamMemberServiceImpl)RemoveTeamMember(ctx context.Context,memberID, teamID string)error{
	return nil
}
func (tm *TeamMemberServiceImpl)GetMemberByTeam(ctx context.Context, teamID string)(*[]model.User,error){
	return nil,nil
}
func (tm *TeamMemberServiceImpl)UpdatePositionTeam(ctx context.Context,update_position int16, id_team,id_member string) error{
	return nil
}
