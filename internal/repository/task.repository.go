package repository

import (
	"context"

	"github.com/StevenAndre/collabtasks/internal/entity"
	"gorm.io/gorm"
)

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) TaskRepository{
	return &GormTaskRepository{
		db: db,
	}
}

func  (r *GormTaskRepository) CreateTask(ctx context.Context,task *entity.Task)error{
	result:=r.db.Create(task)
	return result.Error
}
func  (r *GormTaskRepository) GetAllTasksByProject(ctx context.Context,proj_id string) (*[]entity.Task, error){
	tasks:=&[]entity.Task{}
	result:=r.db.Where("id_project = ?",proj_id).Find(tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks,nil
}
func  (r *GormTaskRepository) GetAllTasksByUser(ctx context.Context,user_id string) (*[]entity.Task, error){
	tasks:=&[]entity.Task{}
	result:=r.db.Where("id_user = ?",user_id).Find(tasks)
	if result.Error!=nil{
		return nil, result.Error
	}
	return tasks,nil
}
func  (r *GormTaskRepository) CompleteTask(ctx context.Context,task_id string) error{
	result:=r.db.Model(&entity.Task{}).Where("id_task = ?", task_id).Update("iscompleted", true)
	if result.Error!=nil{
		return result.Error
	}
	return nil
}
func  (r *GormTaskRepository) GetTasksUncompleted(ctx context.Context,proj_id string) (*[]entity.Task, error){
	tasks:=&[]entity.Task{}
	result:=r.db.Where("id_project = ? AND iscompleted = ?",proj_id,true).Find(tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks,nil
	
}
func  (r *GormTaskRepository) GetTasksCompleted(ctx context.Context,proj_id string) (*[]entity.Task, error){
	tasks:=&[]entity.Task{}
	result:=r.db.Where("id_project = ? AND iscompleted = ?",proj_id,false).Find(tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks,nil
}