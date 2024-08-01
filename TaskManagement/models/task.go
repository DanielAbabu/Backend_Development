package models

type Task struct {
	ID          string `JSON:"id"`
	Title       string `JSON:"title"`
	Description string `JSON:"description"`
	Status      string `JSON:"status"`
}
