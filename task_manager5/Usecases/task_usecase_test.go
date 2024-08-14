package Usecases_test

import (
	"errors"
	"testing"

	"task_manager5/Domain"
	"task_manager5/Usecases"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockTaskRepository is a mock implementation of the TaskRepository interface
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task Domain.Task) (Domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetAllTasks(filter bson.M) ([]Domain.Task, error) {
	args := m.Called(filter)
	return args.Get(0).([]Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) FindTaskById(id string, userId primitive.ObjectID) (Domain.Task, error) {
	args := m.Called(id, userId)
	return args.Get(0).(Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTaskById(id string, task Domain.Task) (Domain.Task, error) {
	args := m.Called(id, task)
	return args.Get(0).(Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) DeleteTaskById(id string, userId primitive.ObjectID) error {
	args := m.Called(id, userId)
	return args.Error(0)
}

func TestPostTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUC := Usecases.NewTaskUC(mockRepo)

	task := Domain.Task{Title: "Test Task"}
	user := Domain.DBUser{ID: primitive.NewObjectID(), Name: "Test User"}

	mockRepo.On("CreateTask", mock.AnythingOfType("Domain.Task")).Return(task, nil)

	createdTask, err := taskUC.PostTask(task, user)

	assert.NoError(t, err)
	assert.Equal(t, task, createdTask)
	mockRepo.AssertExpectations(t)
}

func TestGetTasks(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUC := Usecases.NewTaskUC(mockRepo)

	tasks := []Domain.Task{
		{Title: "Test Task 1"},
		{Title: "Test Task 2"},
	}

	filter := bson.M{"user": "someuser"}
	mockRepo.On("GetAllTasks", filter).Return(tasks, nil)

	result, err := taskUC.GetTasks(filter)

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUC := Usecases.NewTaskUC(mockRepo)

	taskID := "task-id"
	userID := primitive.NewObjectID()
	task := Domain.Task{Title: "Test Task"}

	mockRepo.On("FindTaskById", taskID, userID).Return(task, nil)

	result, err := taskUC.GetTask(taskID, userID)

	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUC := Usecases.NewTaskUC(mockRepo)

	taskID := "task-id"
	task := Domain.Task{Title: "Updated Task"}
	user := Domain.DBUser{ID: primitive.NewObjectID(), Name: "Test User"}

	mockRepo.On("UpdateTaskById", taskID, mock.AnythingOfType("Domain.Task")).Return(task, nil)

	updatedTask, err := taskUC.UpdateTask(taskID, task, user)

	assert.NoError(t, err)
	assert.Equal(t, task, updatedTask)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUC := Usecases.NewTaskUC(mockRepo)

	taskID := "task-id"
	userID := primitive.NewObjectID()

	mockRepo.On("DeleteTaskById", taskID, userID).Return(nil)

	err := taskUC.DeleteTask(taskID, userID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTask_NotFound(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUC := Usecases.NewTaskUC(mockRepo)

	taskID := "non-existent-id"
	userID := primitive.NewObjectID()

	mockRepo.On("FindTaskById", taskID, userID).Return(Domain.Task{}, errors.New("task not found"))

	_, err := taskUC.GetTask(taskID, userID)

	assert.Error(t, err)
	assert.Equal(t, "task not found", err.Error())
	mockRepo.AssertExpectations(t)
}
