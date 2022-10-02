package service

import (
	"net/http"

	"github.com/chen1ting/TravelMaster/internal/models"
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


func (s *Service) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

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

func (s *Service) ValidateToken(c *gin.Context) {
	validateTokenReq := &models.ValidateTokenReq{}
	if err := c.ShouldBindJSON(validateTokenReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validateTokenResp, err := s.server.ValidateToken(c, validateTokenReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, validateTokenResp)
}

func (s *Service) GenerateItinerary(c *gin.Context) {
	generateItineraryReq := &models.GenerateItineraryRequest{}
	if err := c.ShouldBindJSON(generateItineraryReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	generateItineraryResp, err := s.server.GenerateItinerary(c, generateItineraryReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, generateItineraryResp)
}

func (s *Service) GetItinerary(c *gin.Context) {
	getItineraryReq := &models.GetItineraryRequest{}
	if err := c.ShouldBindJSON(getItineraryReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	getItineraryResp, err := s.server.GetItinerary(c, getItineraryReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getItineraryResp)
}

func (s *Service) GetActivitiesByFilter(c *gin.Context) {
	getActivitiesByFilterReq := &models.GetActivitiesByFilterRequest{}
	if err := c.ShouldBindJSON(getActivitiesByFilterReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	getActivitiesByFilterResp, err := s.server.GetActivitiesByFilter(c, getActivitiesByFilterReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getActivitiesByFilterResp)
}