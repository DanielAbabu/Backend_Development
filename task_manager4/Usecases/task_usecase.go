package Usecases

import (
	"task_manager4/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUC struct {
	repo Domain.TaskRepository
}

func NewTaskUC(repository Domain.TaskRepository) *TaskUC {
	return &TaskUC{
		repo: repository,
	}
}

func (TUC *TaskUC) PostTask(task Domain.Task, user Domain.DBUser) (Domain.Task, error) {
	task.User = user
	return TUC.repo.CreateTask(task)
}

func (TUC *TaskUC) GetTasks(filter bson.M) ([]Domain.Task, error) {
	return TUC.repo.GetAllTasks(filter)
}

func (TUC *TaskUC) GetTask(id string, userId primitive.ObjectID) (Domain.Task, error) {
	return TUC.repo.FindTaskById(id, userId)
}

func (TUC *TaskUC) UpdateTask(id string, task Domain.Task, user Domain.DBUser) (Domain.Task, error) {
	task.User = user
	return TUC.repo.UpdateTaskById(id, task)
}

func (TUC *TaskUC) DeleteTask(id string, userId primitive.ObjectID) error {
	return TUC.repo.DeleteTaskById(id, userId)
}
