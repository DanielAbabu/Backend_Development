package models

type User struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password,omitempty" bson:"password"`
	Role      string `json:"role" bson:"role"` // e.g., "admin" or "user"
	UserTasks []Task `json:"usertasks" bson:"usertasks"`
}
