package sqlite

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ritoon/estiam/db"
	"github.com/ritoon/estiam/model"
)

type SQLite struct {
	Conn *gorm.DB
}

func New(dbName string) *db.Storage {
	conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = conn.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	return &db.Storage{
		User: &SQLite{
			Conn: conn,
		},
	}
}

func (c *SQLite) GetByID(id string) (*model.User, error) {
	var u model.User
	err := c.Conn.First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (c *SQLite) GetByEmail(email string) (*model.User, error) {
	var u model.User
	err := c.Conn.First(&u, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (c *SQLite) GetAll() ([]model.User, error) {
	var us []model.User
	err := c.Conn.Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (c *SQLite) DeleteByID(id string) error {
	return c.Conn.Where("id = ?", id).Delete(&model.User{}).Error
}

func (c *SQLite) Create(u *model.User) (*model.User, error) {
	u.ID = uuid.NewString()
	err := c.Conn.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (c *SQLite) Update(id string, data map[string]interface{}) (*model.User, error) {
	u := model.User{ID: id}
	err := c.Conn.Model(&u).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return c.GetByID(id)
}
