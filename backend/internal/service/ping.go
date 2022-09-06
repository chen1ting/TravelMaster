package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
