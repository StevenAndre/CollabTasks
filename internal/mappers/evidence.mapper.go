package mappers

import (
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/model"
)

func (em *EvidenceMapper) FromEntityToModel(evidence *entity.Evidence) *model.Evidence {
	return &model.Evidence{
		EvidenceID: evidence.EvidenceID,
		FilePath: evidence.FilePath,
		Comment: evidence.Comment,
		TaskID: evidence.TaskID,
		CreatedAt: evidence.CreatedAt,
	}
}