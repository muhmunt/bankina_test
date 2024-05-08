package task

import (
	"errors"
	"go-technical-test-bankina/src/entity"
	"go-technical-test-bankina/src/repository"
	"time"
)

type Service interface {
	FindTasks(offset, limit int) ([]entity.Task, error)
	StoreTask(request TaskRequest) (entity.Task, error)
	FindTaskByID(taskID TaskIDRequest) (entity.Task, error)
	UpdateTask(taskID TaskIDRequest, request TaskRequest) (entity.Task, error)
	DeleteTaskByID(taskID TaskIDRequest) (entity.Task, error)
}

type service struct {
	repository repository.TaskRepository
}

func NewService(repository repository.TaskRepository) *service {
	return &service{repository}
}

func (s *service) FindTasks(offset, limit int) ([]entity.Task, error) {
	tasks, err := s.repository.FindAll(offset, limit)

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *service) StoreTask(request TaskRequest) (entity.Task, error) {
	task := entity.Task{}
	task.Title = request.Title
	task.Status = request.Status
	task.Description = request.Description
	task.UserID = request.User.ID
	task.CreatedAt = time.Now()

	newTask, err := s.repository.Save(task)

	if err != nil {
		return newTask, err
	}

	return newTask, nil
}

func (s *service) FindTaskByID(taskID TaskIDRequest) (entity.Task, error) {
	task, err := s.repository.FindByID(taskID.ID)

	if err != nil {
		return task, err
	}

	return task, nil
}

func (s *service) UpdateTask(taskID TaskIDRequest, request TaskRequest) (entity.Task, error) {

	task, err := s.repository.FindByID(taskID.ID)

	if err != nil {
		return task, err
	}

	if task.UserID != request.User.ID {
		return task, errors.New("Unauthorized user update task.")
	}

	task.Title = request.Title
	task.Status = request.Status
	task.Description = request.Description
	task.UpdatedAt = time.Now()

	updateTask, err := s.repository.Save(task)
	if err != nil {
		return updateTask, err
	}

	return updateTask, nil
}

func (s *service) DeleteTaskByID(taskID TaskIDRequest) (entity.Task, error) {
	task, err := s.repository.DeleteTaskByID(taskID.ID)

	if err != nil {
		return task, err
	}

	return task, nil
}
