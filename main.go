package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ritoon/estiam/db/moke"
	"github.com/ritoon/estiam/service"
	"github.com/ritoon/estiam/util"
)

func main() {
	r := gin.Default()
	db := moke.New()
	secretKey := []byte("toto")
	secureJWT := util.MiddlJWT(secretKey)
	s := service.New(db, secretKey)
	r.GET("/users/:id", s.GetUser)
	r.POST("/users", s.CreateUser)
	r.GET("/users", s.GetAllUser)
	r.DELETE("/users/:id", secureJWT, s.DeleteUser)
	r.POST("/login", s.Login)
	r.Run(":8081")
}
