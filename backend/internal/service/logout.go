package service

import (
	"net/http"

	"github.com/chen1ting/TravelMaster/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *Service) LogoutView(c *gin.Context) {
	logoutReq := &models.LogoutReq{}
	if err := c.ShouldBindJSON(logoutReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.server.Logout(c, logoutReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
