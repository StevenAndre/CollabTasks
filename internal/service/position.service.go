package service

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/mappers"
	"github.com/StevenAndre/collabtasks/internal/model"
	"github.com/StevenAndre/collabtasks/internal/repository"
)

var positionMapper mappers.PositionMapper

type PositionServiceImpl struct {
	positionRepo repository.PositionRepository
}

func NewPositionService(posRepo *repository.PositionRepository) PositionsService {
	return &PositionServiceImpl{
		positionRepo: *posRepo,
	}
}

func (ps *PositionServiceImpl) RegisterPosition(ctx context.Context, position dtos.PositionDto) error {
	positionSave := entity.Position{
		Name:        position.Name,
		Description: position.Description,
	}
	err := ps.positionRepo.CreatePosition(ctx, &positionSave)
	return err
}
func (ps *PositionServiceImpl) GetPositionByID(ctx context.Context, positionID int16) (*model.Position, error) {
	position, err := ps.positionRepo.GetPositionById(ctx, positionID)
	if err != nil {
		return nil, err
	}
	return positionMapper.FromEntityToModel(position), nil
}
func (ps *PositionServiceImpl) GetAllPositions(ctx context.Context) (*[]model.Position, error) {
	positionsDB, err := ps.positionRepo.GetAllPositions(ctx)
	if err != nil {
		return nil, err
	}
	positions := []model.Position{}
	for _, value := range *positionsDB {
		positions = append(positions, *positionMapper.FromEntityToModel(&value))
	}

	return &positions, nil
}
func (ps *PositionServiceImpl) UpdatePosition(ctx context.Context, positionID int16, position dtos.PositionDto) error {
	positionDB, err := ps.positionRepo.GetPositionById(ctx, positionID)
	if err != nil {
		return  err
	}
	positionDB.Description=position.Description
	positionDB.Name=position.Name
	err=ps.positionRepo.UpdatePosition(ctx,positionDB)
	return err
}
func (ps *PositionServiceImpl) RemovePosition(ctx context.Context, positionID int16) error {
	return ps.positionRepo.DeletePosition(ctx, positionID)
}
