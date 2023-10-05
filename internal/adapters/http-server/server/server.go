package server

import (
	"time"

	"github.com/gin-gonic/gin"
)

type handler interface {
	AddUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type middleware interface {
	// Logger(c *gin.Context)
	RequestID(c *gin.Context)
}

type config interface {
	GetHttpPort() string
	GetTimeLimit() time.Duration
}

type server struct {
	port      string
	timeLimit time.Duration
	s         *gin.Engine
	handler   handler
}

func New(cfg config, h handler, m middleware) *server {

	server := &server{
		port:      cfg.GetHttpPort(),
		timeLimit: cfg.GetTimeLimit(),
		handler:   h,
	}

	s := gin.New()

	s.Use(m.RequestID)
	s.Use(gin.Logger())

	server.s = s
	server.bind()

	return server
}

func (s server) bind() {
	s.s.POST("/user", s.handler.AddUser)
	s.s.GET("/user", s.handler.GetUsers)
	s.s.GET("/user/id", s.handler.GetUserById)
	s.s.DELETE("/user", s.handler.DeleteUser)
	s.s.PATCH("/user", s.handler.UpdateUser)
}

func (s server) Run() error {
	return s.s.Run(":" + s.port)
}
