package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Status      string             `json:"status" bson:"status"`
	DueDate     time.Time          `json:"due_date" bson:"due_date"`
	User        DBUser             `json:"user"`
}

type TaskUsecase interface {
	PostTask(Task, DBUser) (Task, error)
	GetTasks(bson.M) ([]Task, error)
	GetTask(string, primitive.ObjectID) (Task, error)
	UpdateTask(string, Task, DBUser) (Task, error)
	DeleteTask(string, primitive.ObjectID) error
}

type TaskRepository interface {
	CreateTask(Task) (Task, error)
	DeleteTaskById(string, primitive.ObjectID) error
	UpdateTaskById(string, Task) (Task, error)
	GetAllTasks(bson.M) ([]Task, error)
	FindTaskById(string, primitive.ObjectID) (Task, error)
}
