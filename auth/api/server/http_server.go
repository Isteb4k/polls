package server

import (
	"auth/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server interface {
	Run() error
}

type server struct {
	router *gin.Engine
	users  db.Users
}

func New(users db.Users) Server {
	router := gin.Default()

	s := server{
		router: router,
		users:  users,
	}

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user": "dsa"})
	})
	router.POST("/create_user", s.createUserHandler)
	router.DELETE("/delete_user/:id", s.deleteUserHandler)
	router.GET("/get_user/:id", s.getUserHandler)

	return &s
}

func (s *server) Run() error {
	return s.router.Run(":8081")
}
