package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormPositionRepository struct {
	db *gorm.DB
}

func NewGormPositionRepository(db *gorm.DB) PositionRepository {
	return &GormPositionRepository{db: db}
}

func (por *GormPositionRepository) CreatePosition(ctx context.Context, position *entity.Position) error {
	result := por.db.Create(position)
	return result.Error
}

func (por *GormPositionRepository) GetPositionById(ctx context.Context, position_id int16) (*entity.Position, error) {
	position := &entity.Position{
		PositionID: position_id,
	}
	result := por.db.First(position)
	if result.Error != nil {
		return nil, result.Error
	}
	return position, nil
}

func (por *GormPositionRepository) GetAllPositions(ctx context.Context) (*[]entity.Position, error) {
	positions := &[]entity.Position{}
	result := por.db.Find(positions)
	if result.Error != nil {
		return nil, result.Error
	}
	return positions, nil
}

func (por *GormPositionRepository) UpdatePosition(ctx context.Context, newPosition *entity.Position) error {
	result := por.db.Save(newPosition)
	return result.Error
}

func (por *GormPositionRepository) DeletePosition(ctx context.Context, position_id int16) error {
	position := &entity.Position{
		PositionID: position_id,
	}
	result := por.db.Delete(position)
	return result.Error
}
