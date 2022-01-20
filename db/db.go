package db

import "github.com/ritoon/estiam/model"

type Storage struct {
	User StorageUser
}

type StorageUser interface {
	GetByID(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]model.User, error)
	DeleteByID(id string) error
	Create(u *model.User) (*model.User, error)
	Update(id string, data map[string]interface{}) (*model.User, error)
}
