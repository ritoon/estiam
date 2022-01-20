package moke

import (
	"errors"

	"github.com/google/uuid"

	"github.com/ritoon/estiam/db"
	"github.com/ritoon/estiam/model"
)

var _ db.Storage = &Moke{}

type Moke struct {
	listUser map[string]*model.User
}

func New() *Moke {
	return &Moke{
		listUser: make(map[string]*model.User),
	}
}

func (m *Moke) GetUserByID(id string) (*model.User, error) {
	u, ok := m.listUser[id]
	if !ok {
		return nil, errors.New("db user: not found")
	}
	return u, nil
}

func (m *Moke) DeleteUserByID(id string) error {
	_, ok := m.listUser[id]
	if !ok {
		return errors.New("db user: not found")
	}
	delete(m.listUser, id)
	return nil
}

func (m *Moke) CreateUser(u *model.User) (*model.User, error) {
	u.ID = uuid.New().String()
	m.listUser[u.ID] = u
	return u, nil
}

func (m *Moke) UpdateUser(id string, data map[string]interface{}) (*model.User, error) {
	u, ok := m.listUser[id]
	if !ok {
		return nil, errors.New("db user: not found")
	}
	if value, ok := data["first_name"]; ok {
		u.FirstName = value.(string)
	}
	if value, ok := data["last_name"]; ok {
		u.FirstName = value.(string)
	}
	return nil, nil
}

func (m *Moke) GetAllUser() ([]model.User, error) {
	us := make([]model.User, len(m.listUser))
	var i int
	for k := range m.listUser {
		if m.listUser[k] != nil {
			us[i] = *m.listUser[k]
		}
		i++
	}
	return us, nil
}
