package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()
	r.GET("/users/:id", HandlerGetUser)
	r.POST("/users", HandlerCreateUser)
	r.GET("/users", HandlerGetAllUser)
	r.DELETE("/users/:id", HandlerDeleteUser)
	r.Run(":8081")
}

var listUser map[string]*User = map[string]*User{
	"abcd": {FirstName: "Bob", LastName: "Pike"},
}

type User struct {
	ID        string `json:"id"`
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

func HandlerGetAllUser(c *gin.Context) {
	if len(listUser) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("no users found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": listUser,
	})
}

func HandlerCreateUser(c *gin.Context) {
	var u User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	u.ID = uuid.New().String()
	listUser[u.ID] = &u

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func HandlerDeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, ok := listUser[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	delete(listUser, id)
	c.JSON(http.StatusAccepted, gin.H{
		"delete": id,
	})
}
