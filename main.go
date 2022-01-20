package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ritoon/estiam/db/moke"
	"github.com/ritoon/estiam/service"
)

func main() {
	r := gin.Default()
	db := moke.New()
	s := service.New(db)
	r.GET("/users/:id", s.GetUser)
	r.POST("/users", s.CreateUser)
	r.GET("/users", s.GetAllUser)
	r.DELETE("/users/:id", s.DeleteUser)
	r.POST("/login", s.Login)
	r.Run(":8081")
}
