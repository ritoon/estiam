package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users/:id", HandlerGetUser)
	// TODO: implement all routes
	//r.POST("/users",....)
	//r.GET("/users",....)
	//r.DELETE("/users",....)

	r.Run(":8081")
}

var listUser map[string]*User = map[string]*User{
	"abcd": {FirstName: "Bob", LastName: "Pike"},
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func HandlerGetUser(c *gin.Context) {
	id := c.Param("id")
	u, ok := listUser[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
