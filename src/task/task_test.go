package task

import (
	"errors"
	"go-technical-test-bankina/src/entity"
	"testing"
	"time"
)

type MockTaskRepository struct {
	data map[int]entity.Task
}

func (m *MockTaskRepository) FindAll(offset, limit int) ([]entity.Task, error) {
	var tasks []entity.Task
	for _, task := range m.data {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (m *MockTaskRepository) Save(task entity.Task) (entity.Task, error) {
	if m.data == nil {
		m.data = make(map[int]entity.Task)
	}
	m.data[task.ID] = task
	return task, nil
}

func (m *MockTaskRepository) FindByID(taskID int) (entity.Task, error) {
	if task, ok := m.data[taskID]; ok {
		return task, nil
	}
	return entity.Task{}, errors.New("task not found")
}

func (m *MockTaskRepository) DeleteTaskByID(taskID int) (entity.Task, error) {
	task, ok := m.data[taskID]
	if !ok {
		return entity.Task{}, errors.New("task not found")
	}
	delete(m.data, taskID)
	return task, nil
}

func TestFindTasks(t *testing.T) {
	mockRepo := &MockTaskRepository{}

	service := NewService(mockRepo)

	testTasks := []entity.Task{
		{ID: 1, Title: "Task 1", Status: "pending", UserID: 1, CreatedAt: time.Now()},
		{ID: 2, Title: "Task 2", Status: "pending", UserID: 1, CreatedAt: time.Now()},
		{ID: 3, Title: "Task 3", Status: "pending", UserID: 1, CreatedAt: time.Now()},
	}
	mockRepo.data = make(map[int]entity.Task)
	for _, task := range testTasks {
		mockRepo.data[task.ID] = task
	}

	// Test FindTasks method
	tasks, err := service.FindTasks(0, 10)
	if err != nil {
		t.Errorf("Error getting tasks: %v", err)
	}
	if len(tasks) != len(testTasks) {
		t.Errorf("Expected %d tasks, got %d", len(testTasks), len(tasks))
	}
}

func TestStoreTask(t *testing.T) {
	mockRepo := &MockTaskRepository{}

	service := NewService(mockRepo)

	request := TaskRequest{
		Title:       "Test Task",
		Status:      "pending",
		Description: "Test Description",
		User: entity.User{
			ID: 1,
		},
	}

	task, err := service.StoreTask(request)
	if err != nil {
		t.Errorf("Error storing task: %v", err)
	}

	// Check if task is saved in the repository
	_, err = mockRepo.FindByID(task.ID)
	if err != nil {
		t.Errorf("Stored task not found in the repository")
	}
}

func TestFindTaskByID(t *testing.T) {
	mockRepo := &MockTaskRepository{}

	service := NewService(mockRepo)

	testTask := entity.Task{
		ID:          1,
		Title:       "Test Task",
		Status:      "pending",
		Description: "Test Description",
		UserID:      1,
		CreatedAt:   time.Now(),
	}
	mockRepo.data = map[int]entity.Task{
		1: testTask,
	}

	task, err := service.FindTaskByID(TaskIDRequest{ID: 1})
	if err != nil {
		t.Errorf("Error finding task by ID: %v", err)
	}
	if task.ID != testTask.ID {
		t.Errorf("Expected task ID %d, got %d", testTask.ID, task.ID)
	}
}

func TestUpdateTask(t *testing.T) {
	mockRepo := &MockTaskRepository{}

	service := NewService(mockRepo)

	testTask := entity.Task{
		ID:          1,
		Title:       "Test Task",
		Status:      "pending",
		Description: "Test Description",
		UserID:      1,
		CreatedAt:   time.Now(),
	}
	mockRepo.data = map[int]entity.Task{
		1: testTask,
	}

	// Test UpdateTask method with correct user ID
	updateRequest := TaskRequest{
		Title:       "Updated Task",
		Status:      "done",
		Description: "Updated Description",
		User: entity.User{
			ID: 1,
		},
	}
	updatedTask, err := service.UpdateTask(TaskIDRequest{ID: 1}, updateRequest)
	if err != nil {
		t.Errorf("Error updating task: %v", err)
	}
	if updatedTask.Title != updateRequest.Title {
		t.Errorf("Expected updated task name %s, got %s", updateRequest.Title, updatedTask.Title)
	}

	// Test UpdateTask method with incorrect user ID
	updateRequest.User.ID = 2
	_, err = service.UpdateTask(TaskIDRequest{ID: 1}, updateRequest)
	if err == nil {
		t.Errorf("Expected error for unauthorized user update")
	}
}

func TestDeleteTaskByID(t *testing.T) {
	mockRepo := &MockTaskRepository{}

	service := NewService(mockRepo)

	testTask := entity.Task{
		ID:          1,
		Title:       "Test task",
		Status:      "pending",
		Description: "Test Description",
		UserID:      1,
		CreatedAt:   time.Now(),
	}
	mockRepo.data = map[int]entity.Task{
		1: testTask,
	}

	deletedTask, err := service.DeleteTaskByID(TaskIDRequest{ID: 1})
	if err != nil {
		t.Errorf("Error deleting task by ID: %v", err)
	}
	if deletedTask.ID != testTask.ID {
		t.Errorf("Expected deleted task ID %d, got %d", testTask.ID, deletedTask.ID)
	}
}
