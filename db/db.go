package db

import "github.com/ritoon/estiam/model"

type Storage interface {
	GetUserByID(id string) (*model.User, error)
	GetAllUser() ([]model.User, error)
	DeleteUserByID(id string) error
	CreateUser(u *model.User) (*model.User, error)
	UpdateUser(id string, data map[string]interface{}) (*model.User, error)
}
