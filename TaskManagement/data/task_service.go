package data

import (
	"TaskManagement/models"
	"errors"
	"strconv"
)

var tasks map[string]models.Task = map[string]models.Task{
	"1": {
		ID:          "1",
		Title:       "Task 1",
		Description: "A real-time chat feature using WebSockets",
		Status:      "NOt Started",
	},
	"2": {
		ID:          "2",
		Title:       "Task 2",
		Description: "Create a RESTful API for a blogging platform",
		Status:      "In Progress",
	},
	"3": {
		ID:          "3",
		Title:       "Task 3",
		Description: "Design a user authentication system with JWT",
		Status:      "Not Started",
	},
	"4": {
		ID:          "4",
		Title:       "Task 4",
		Description: "Implement a real-time chat feature using WebSockets",
		Status:      "Complete",
	},
	"5": {
		ID:          "5",
		Title:       "Task 5",
		Description: "Build a responsive front-end with React and Tailwind CSS",
		Status:      "Complete",
	},
}

var TaskIDs = 6

func Add(new models.Task) models.Task {
	new.ID = strconv.Itoa(TaskIDs)
	TaskIDs += 1
	tasks[new.ID] = new
	return new
}

func Update(ID string, new models.Task) error {

	if _, err := tasks[ID]; !err {
		return errors.New("task not found")
	}
	new.ID = strconv.Itoa(TaskIDs)

	tasks[ID] = new
	return nil
}

func Delete(ID string) error {
	if _, err := tasks[ID]; !err {
		return errors.New("task not found")
	}
	delete(tasks, ID)
	return nil
}

func GetAll() map[string]models.Task {
	return tasks
}

func Get(id string) (models.Task, error) {
	task, err := tasks[id]
	if !err {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}
