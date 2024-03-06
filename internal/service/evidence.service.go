package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/mappers"
	"github.com/StevenAndre/collabtasks/internal/model"
	"github.com/StevenAndre/collabtasks/internal/repository"
	"github.com/StevenAndre/collabtasks/utils"
)
var evidenceMapper mappers.EvidenceMapper

type EvidenceServiceImpl struct {
	evidenceRepo repository.EvidenceRepository
	taskRepo repository.TaskRepository
}

func NewEvidenceService(er *repository.EvidenceRepository, tr  *repository.TaskRepository) EvidenceService{
	return &EvidenceServiceImpl{
		evidenceRepo:*er,
		taskRepo: *tr,
	}
}

func(es *EvidenceServiceImpl) RegisterEvidence(ctx context.Context,evidenceDto dtos.NotificationDto,taskID string, file *multipart.FileHeader) error{
	task,err:=es.taskRepo.GetTaskByID(ctx,taskID)
	if err!=nil{
		return err
	}
	var evidencepath string
	if file == nil {
		evidencepath = ""
	} else {
		evidencepath, err = saveFile(file)
		if err != nil {
			return err
		}
	}
	id:=utils.GeneradorIds(10)
	evidence:=entity.Evidence{
		EvidenceID: id,
		Comment: evidenceDto.Message,
		CreatedAt: time.Now(),
		TaskID: task.TaskID,
		FilePath: evidencepath ,
	}

	err=es.evidenceRepo.SaveEvidence(ctx,&evidence)
	return err
}

func(es *EvidenceServiceImpl) ShowEvidenceByID(ctx context.Context,evidenceID string)(*model.Evidence,error){
	evidenceDB,err:=es.evidenceRepo.GetEvidenceById(ctx,evidenceID)
	if err!=nil{
		return nil,err
	}
	evidence:=evidenceMapper.FromEntityToModel(evidenceDB)
	return evidence,nil
}
func(es *EvidenceServiceImpl) ViewEvidencesByTask(ctx context.Context,taskID string)(*[]model.Evidence,error){
	evidencesList,err:=es.evidenceRepo.GetEvidencesByTask(ctx,taskID)
	if err!=nil{
		return nil,err
	}
	evidences:=[]model.Evidence{}
	for _,v:=range *evidencesList{
		evidences=append(evidences,*evidenceMapper.FromEntityToModel(&v))
	}

	return &evidences,nil
}
func(es *EvidenceServiceImpl) DeleteEvidence(ctx context.Context,evidenceID string) error{
	err:=es.evidenceRepo.DeleteEvidence(ctx,evidenceID)
	return err
}