package Domain

import (
	"time"

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
	GetTasks(userid string) ([]Task, error)
	GetTask(string, string) (Task, error)
	UpdateTask(string, Task, DBUser) (Task, error)
	DeleteTask(string, string) error
}

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks(userid string) ([]Task, error)
	FindTaskById(id string, userId string) (Task, error)
	UpdateTaskById(id string, task Task) (Task, error)
	DeleteTaskById(id string, userId string) error
}
