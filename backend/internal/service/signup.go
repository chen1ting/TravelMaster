package service

import (
	"net/http"

	"github.com/chen1ting/TravelMaster/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *Service) SignupView(c *gin.Context) {
	signupReq := &models.SignupReq{}
	if err := c.ShouldBindJSON(&signupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	signupResp, err := s.server.Signup(c, signupReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, signupResp)
}
