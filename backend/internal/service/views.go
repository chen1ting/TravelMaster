package service

import (
	"github.com/chen1ting/TravelMaster/internal/server"
	"github.com/gin-gonic/gin"
)

type ServiceInf interface {
	Ping(c *gin.Context)
	SignupView(c *gin.Context)
}

type Service struct {
	server *server.Server
}

var _ ServiceInf = (*Service)(nil)

func NewService() *Service {
	return &Service{
		server: server.NewServer(),
	}
}
