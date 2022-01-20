package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ritoon/estiam/service"
)

func main() {
	r := gin.Default()
	r.GET("/users/:id", service.HandlerGetUser)
	r.POST("/users", service.HandlerCreateUser)
	r.GET("/users", service.HandlerGetAllUser)
	r.DELETE("/users/:id", service.HandlerDeleteUser)
	r.Run(":8081")
}
