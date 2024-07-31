package models

type Task struct {
	ID      string `JSON:"id"`
	Name    string `JSON:"name"`
	Descrip string `JSON:"desc"`
	Status  string `JSON:"status"`
}
