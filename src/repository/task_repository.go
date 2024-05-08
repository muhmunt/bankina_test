package repository

import (
	"go-technical-test-bankina/src/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Save(task entity.Task) (entity.Task, error)
	FindAll(offset, limit int) ([]entity.Task, error)
	FindByID(taskID int) (entity.Task, error)
	DeleteTaskByID(task int) (entity.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTask(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Save(task entity.Task) (entity.Task, error) {
	err := r.db.Save(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *taskRepository) FindAll(offset, limit int) ([]entity.Task, error) {
	var tasks []entity.Task

	err := r.db.Offset(offset).Limit(limit).Find(&tasks).Error

	if err != nil {
		return tasks, err
	}

	return tasks, nil

}

func (r *taskRepository) FindByID(taskID int) (entity.Task, error) {
	var getTask entity.Task

	err := r.db.Where("id = ?", taskID).First(&getTask).Error

	if err != nil {
		return getTask, err
	}

	return getTask, nil

}

func (r *taskRepository) DeleteTaskByID(task int) (entity.Task, error) {
	var deleteTask entity.Task
	err := r.db.Where("id = ?", task).Delete(&deleteTask).Error

	if err != nil {
		return deleteTask, err
	}

	return deleteTask, nil

}
