package service

import (
	"errors"
	"net/http"

	"github.com/dgkg/cmi/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ritoon/estiam/db/moke"
	"github.com/ritoon/estiam/model"
)

type Service struct {
	db db.Storage
}

// go to Service folder.
func (s *Service) HandlerGetUser(c *gin.Context) {
	id := c.Param("id")
	u, ok := s.db.GetUserByID(id)
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
	if len(moke.ListUser) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("no users found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": moke.ListUser,
	})
}

func HandlerCreateUser(c *gin.Context) {
	var u model.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	u.ID = uuid.New().String()
	moke.ListUser[u.ID] = &u

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func HandlerDeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, ok := moke.ListUser[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	delete(moke.ListUser, id)
	c.JSON(http.StatusAccepted, gin.H{
		"delete": id,
	})
}
