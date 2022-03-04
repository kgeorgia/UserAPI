package repository

import "refactoring/internal/model"

type Storage interface {
	SearchUsers() *UserList
	GetByID(string) (User, error)
	CreateUser(model.CreateUserRequest) (string, error)
	UpdateUser(string, model.UpdateUserRequest) error
	DeleteUser(string) error
}
