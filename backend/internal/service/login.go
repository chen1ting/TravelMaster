package service

import (
	"net/http"

	"github.com/chen1ting/TravelMaster/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *Service) LoginView(c *gin.Context) {
	loginReq := &models.LoginReq{}
	if err := c.ShouldBindJSON(loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loginResp, err := s.server.Login(c, loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loginResp)
}
