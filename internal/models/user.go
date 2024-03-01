package models

import "github.com/google/uuid"

type UserRepository interface {
	CreateUser(u *User) error
	ListAll() ([]*User, error)
}

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Age    int64  `json:"age"`
}

func NewUser(name, status string, age int64) *User {
	return &User{
		ID:     uuid.New().String(),
		Name:   name,
		Status: status,
		Age:    age,
	}
}
