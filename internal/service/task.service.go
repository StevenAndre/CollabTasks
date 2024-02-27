package service

import (
	"context"
	"time"

	"github.com/StevenAndre/collabtasks/internal/api/dtos"
	"github.com/StevenAndre/collabtasks/internal/entity"
	"github.com/StevenAndre/collabtasks/internal/mappers"
	"github.com/StevenAndre/collabtasks/internal/model"
	"github.com/StevenAndre/collabtasks/internal/repository"
	"github.com/StevenAndre/collabtasks/utils"
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
	user, err := ts.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	project, err := ts.projectRepo.GetProjectByID(ctx, projectId)
	if err != nil {
		return err
	}
	id := utils.GeneradorIds(10)
	task := &entity.Task{
		TaskID:      id,
		Title:       taskDto.Title,
		Description: taskDto.Description,
		CreatedAt:   time.Now(),
		IsCompleted: false,
		UserID:      user.UserID,
		ProjectID:   project.ProjectID,
	}
	err = ts.taskRepo.CreateTask(ctx, task)
	return err
}
func (ts *TaskServiceImpl) ListTasksByproject(ctx context.Context, projectId string) (*[]model.TaskU, error) {
	project, err := ts.projectRepo.GetProjectByID(ctx, projectId)
	if err != nil {
		return nil, err
	}
	tasksDB, err := ts.taskRepo.GetAllTasksByProject(ctx, project.ProjectID)
	if err != nil {
		return nil, err
	}
	tasks := []model.TaskU{}
	for _, v := range *tasksDB {
		tasks = append(tasks, *taskMapper.FromEntitytoModelU(&v))
	}

	return &tasks, nil
}
func (ts *TaskServiceImpl) ListTasksByUser(ctx context.Context, userID string) (*[]model.Task, error) {
	user, err := ts.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil,err
	}
	tasksDB, err := ts.taskRepo.GetAllTasksByUser(ctx,user.UserID)
	if err!=nil{
		return nil,err
	}
	tasks := []model.Task{}
	for _, v := range *tasksDB {
		tasks = append(tasks, *taskMapper.FromEntitytoModel(&v))
	}
	return &tasks, nil
}
func (ts *TaskServiceImpl) CompleteTask(ctx context.Context, taskId string) error {
	t, err := ts.taskRepo.GetTaskByID(ctx, taskId)
	if err!=nil{
		return err
	}
	err=ts.taskRepo.CompleteTask(ctx,t.TaskID)
	return err
}
func (ts *TaskServiceImpl) ListTasksUncompleted(ctx context.Context, projectId string) (*[]model.TaskU, error) {
	project, err := ts.projectRepo.GetProjectByID(ctx, projectId)
	if err != nil {
		return nil, err
	}
	tasksDB, err := ts.taskRepo.GetTasksUncompleted(ctx, project.ProjectID)
	if err != nil {
		return nil, err
	}
	tasks := []model.TaskU{}
	for _, v := range *tasksDB {
		tasks = append(tasks, *taskMapper.FromEntitytoModelU(&v))
	}

	return &tasks, nil
}
func (ts *TaskServiceImpl) ListTasksCompleted(ctx context.Context, projectId string) (*[]model.TaskU, error) {
	project, err := ts.projectRepo.GetProjectByID(ctx, projectId)
	if err != nil {
		return nil, err
	}
	tasksDB, err := ts.taskRepo.GetTasksCompleted(ctx, project.ProjectID)
	if err != nil {
		return nil, err
	}
	tasks := []model.TaskU{}
	for _, v := range *tasksDB {
		tasks = append(tasks, *taskMapper.FromEntitytoModelU(&v))
	}

	return &tasks, nil
}
