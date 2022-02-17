package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ritoon/estiam/db"
	"github.com/ritoon/estiam/db/sqlite"
	"github.com/ritoon/estiam/model"
)

type MySQL = sqlite.SQLite

func New(dbName, user, pass, port string) *db.Storage {
	// dsn := "estiam-user:estiam-pwd@tcp(localhost:3306)/estiam-db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(localhost:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, port, dbName)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = conn.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	return &db.Storage{
		User: &MySQL{
			Conn: conn,
		},
	}
}
