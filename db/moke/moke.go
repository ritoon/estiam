package moke

import (
	"errors"

	"github.com/ritoon/estiam/model"
)

type Moke struct {
	listUser map[string]*model.User
}

func (m *Moke) GetUserByID(id string) (*model.User, error) {
	u, ok := m.listUser[id]
	if !ok {
		return nil, errors.New("db user: not found")
	}
	return u, nil
}

func (m *Moke) DeleteUserByID(id string) error {
	return nil
}
func (m *Moke) CreateUser(u *model.User) (*model.User, error) {
	return nil, nil
}
func (m *Moke) UpdateUser(id string, data map[string]interface{}) (*model.User, error) {
	return nil, nil
}
