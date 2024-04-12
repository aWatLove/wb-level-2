package model

// User - модель структуры пользователя
type User struct {
	ID     string           `json:"id"`
	Events map[string]Event `json:"events"`
}

// NewUser - конструктор модели User
func NewUser(id string) User {
	return User{ID: id, Events: make(map[string]Event)}
}
