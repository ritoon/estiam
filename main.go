package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/ritoon/estiam/db"
	"github.com/ritoon/estiam/db/moke"
	"github.com/ritoon/estiam/db/mysql"
	"github.com/ritoon/estiam/db/sqlite"
	"github.com/ritoon/estiam/service"
	"github.com/ritoon/estiam/util"
)

type Config struct {
	EnvType    string
	ListenPort string
	SecretKey  []byte
	db         struct {
		DBName string
		User   string
		Pass   string
		Port   string
	}
}

var config Config

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	config.EnvType = viper.GetString("EnvType")
	config.SecretKey = []byte(viper.GetString("SecretKey"))
	config.ListenPort = viper.GetString("ListenPort")
	// connect to DB
	config.db.DBName = viper.GetString("db.DBName")
	config.db.User = viper.GetString("db.User")
	config.db.Pass = viper.GetString("db.Pass")
	config.db.Port = viper.GetString("db.Port")
}

func main() {
	r := gin.Default()
	var db *db.Storage
	log.Println("ENV:", config.EnvType)
	if config.EnvType == "dev" {
		log.Println("create Moke DB")
		db = moke.New()

	} else if config.EnvType == "prod" {
		log.Println("connect to an MySQL")
		db = mysql.New(config.db.DBName, config.db.User, config.db.Pass, config.db.Port)
	} else {
		log.Println("create SQLite DB")
		db = sqlite.New("storage.db")
	}

	secureJWT := util.MiddlJWT(config.SecretKey)
	s := service.New(db, config.SecretKey)
	r.GET("/users/:id", s.GetUser)
	r.POST("/users", s.CreateUser)
	r.GET("/users", s.GetAllUser)
	r.DELETE("/users/:id", secureJWT, s.DeleteUser)
	r.POST("/login", s.Login)
	r.Run(":" + config.ListenPort)
}
