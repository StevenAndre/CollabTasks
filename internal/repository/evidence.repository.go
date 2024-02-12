package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormEvidenceRepository struct {
	db *gorm.DB
}

func NewGormEvidenceRepository(db *gorm.DB) EvidenceRepository{
	return &GormEvidenceRepository{
		db: db,
	}
}

func (er *GormEvidenceRepository) SaveEvidence(ctx context.Context, evidence *entity.Evidence)error{
	result:=er.db.Create(evidence)
	return result.Error
}
func (er *GormEvidenceRepository) GetEvidenceById(ctx context.Context, evidence_id string)(*entity.Evidence,error){
	evidence:=&entity.Evidence{
		EvidenceID: evidence_id,
	}
	result:=er.db.First(evidence)
	if result.Error!=nil{
		return nil,result.Error
	}
	return evidence,nil
}

func (er *GormEvidenceRepository) GetEvidencesByTask(ctx context.Context, task_id string)(*[]entity.Evidence,error){
	evidences:=[]entity.Evidence{}
	result:=er.db.Where("task_id = ?",task_id).Find(&evidences)
	if result.Error!=nil{
		return nil,result.Error
	}
	return &evidences,nil
}

func (er *GormEvidenceRepository) DeleteEvidence(ctx context.Context,evidence_id string)error{
	result:=er.db.Delete(&entity.Evidence{
		EvidenceID: evidence_id,
	})
	return result.Error
}
