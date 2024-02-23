package service

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/StevenAndre/collabtasks/internal/mappers"
	"github.com/StevenAndre/collabtasks/internal/model"
	"github.com/StevenAndre/collabtasks/internal/repository"
)

var taskMapper mappers.TaskMapper

type TaskServiceImpl struct {
	taskRepo    repository.TaskRepository
	userRepo    repository.UserRepository
	projectRepo repository.ProjectRepository
}

func NewTaskService(tr *repository.TaskRepository,
	ur *repository.UserRepository,
	pr *repository.ProjectRepository) TaskService {
	return &TaskServiceImpl{*tr, *ur, *pr}
}

func (ts *TaskServiceImpl) RegisterTask(ctx context.Context, taskDto dtos.TaskDto, userID, projectId string) error {
	return nil
}
func (ts *TaskServiceImpl) ListTasksByproject(ctx context.Context, projectId string) (*[]model.TaskU, error) {
	return nil, nil
}
func (ts *TaskServiceImpl) ListTasksByUser(ctx context.Context, projectId string) (*[]model.Task, error) {
	return nil, nil
}
func (ts *TaskServiceImpl) CompleteTask(ctx context.Context, taskId string) error {
	return nil
}
func (ts *TaskServiceImpl) ListTasksUncompleted(ctx context.Context, projectId string) (*[]model.TaskU, error) {
	return nil, nil
}
func (ts *TaskServiceImpl) ListTasksCompleted(ctx context.Context, projectId string) (*[]model.TaskU, error) {
	return nil, nil
}
