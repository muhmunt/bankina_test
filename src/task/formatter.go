package task

import (
	"go-technical-test-bankina/src/entity"
	"time"
)

type TaskFormatter struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FormatTask(task entity.Task) TaskFormatter {
	formatter := TaskFormatter{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	return formatter
}

func FormatTasks(tasks []entity.Task) []TaskFormatter {
	tasksFormatter := []TaskFormatter{}

	for _, task := range tasks {
		TaskFormatter := FormatTask(task)
		tasksFormatter = append(tasksFormatter, TaskFormatter)
	}

	return tasksFormatter
}
