package service

import (
	"github.com/chen1ting/TravelMaster/internal/models"
	"github.com/chen1ting/TravelMaster/internal/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServiceInf interface {
	Ping(c *gin.Context)
	SignupView(c *gin.Context)
	// TODO: add in all endpoints here?
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

func (s *Service) CreateActivity(c *gin.Context) {
	createActivityForm := &models.CreateActivityForm{}
	if err := c.ShouldBind(&createActivityForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createActivityResp, err := s.server.CreateActivity(createActivityForm, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createActivityResp)
}

func (s *Service) UpdateActivity(c *gin.Context) {
	updateActivityForm := &models.UpdateActivityForm{}
	if err := c.ShouldBind(&updateActivityForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateActivityResp, err := s.server.UpdateActivity(updateActivityForm, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, updateActivityResp)
}

func (s *Service) GetActivity(c *gin.Context) {
	getActivityReq := &models.GetActivityReq{}
	if err := c.ShouldBindJSON(&getActivityReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	getActivityResp, err := s.server.GetActivity(getActivityReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getActivityResp)
}

func (s *Service) SearchActivity(c *gin.Context) {
	searchActivityReq := &models.SearchActivityReq{}
	if err := c.ShouldBindJSON(&searchActivityReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	searchActivityResp, err := s.server.SearchActivity(searchActivityReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, searchActivityResp)
}

func (s *Service) ReportInactiveActivity(c *gin.Context) {
	inactivateActivityReq := &models.InactivateActivityReq{}
	if err := c.ShouldBindJSON(&inactivateActivityReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inactivateActivityResp, err := s.server.ReportInactiveActivity(inactivateActivityReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, inactivateActivityResp)
}

func (s *Service) DeleteActivityImage(c *gin.Context) {
	deleteActivityImageReq := &models.DeleteActivityImageReq{}
	if err := c.ShouldBindJSON(&deleteActivityImageReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deleteActivityImageResp, err := s.server.DeleteActivityImage(deleteActivityImageReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, deleteActivityImageResp)
}
