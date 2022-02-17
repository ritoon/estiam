package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/ritoon/estiam/db/moke"
	"github.com/ritoon/estiam/service"
	"github.com/ritoon/estiam/util"
)

type Config struct {
	ListenPort string
	SecretKey  []byte
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
	config.SecretKey = []byte(viper.GetString("SecretKey"))
	config.ListenPort = viper.GetString("ListenPort")
}

func main() {
	r := gin.Default()
	db := moke.New()
	secureJWT := util.MiddlJWT(config.SecretKey)
	s := service.New(db, config.SecretKey)
	r.GET("/users/:id", s.GetUser)
	r.POST("/users", s.CreateUser)
	r.GET("/users", s.GetAllUser)
	r.DELETE("/users/:id", secureJWT, s.DeleteUser)
	r.POST("/login", s.Login)
	r.Run(":" + config.ListenPort)
}
