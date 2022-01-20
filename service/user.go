package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ritoon/estiam/db"
	"github.com/ritoon/estiam/model"
)

type Service struct {
	db db.Storage
}

func New(db db.Storage) *Service {
	return &Service{
		db: db,
	}
}

// go to Service folder.
func (s *Service) GetUser(c *gin.Context) {
	id := c.Param("id")
	u, err := s.db.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func (s *Service) GetAllUser(c *gin.Context) {
	us, err := s.db.GetAllUser()
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": us,
	})
}

func (s *Service) CreateUser(c *gin.Context) {
	var u model.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	_, err = s.db.CreateUser(&u)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func (s *Service) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error id": id,
		})
		return
	}
	err := s.db.DeleteUserByID(id)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"delete": id,
	})
}
