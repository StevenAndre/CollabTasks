package mappers

import (
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/model"
)
var userMapper UserMapper
func (tm *TaskMapper) FromEntitytoModel(task *entity.Task) *model.Task {
	return &model.Task{
		TaskID:      task.TaskID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt,
		IsCompleted: task.IsCompleted,
		UserID:      task.UserID,
		ProjectID:   task.ProjectID,
	}
}
func (tm *TaskMapper) FromEntitytoModelU(task *entity.Task) *model.TaskU {
	return &model.TaskU{
		TaskID:      task.TaskID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt,
		IsCompleted: task.IsCompleted,
		UserID:      task.UserID,
		ProjectID:   task.ProjectID,
		User:        *userMapper.FromEntityToModel(&task.User),
	}
}
