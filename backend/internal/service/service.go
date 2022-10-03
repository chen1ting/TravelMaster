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
	signupForm := &models.SignupForm{}
	if err := c.ShouldBind(&signupForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	signupResp, err := s.server.Signup(c, signupForm)
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

func (s *Service) UpdateItinerary(c *gin.Context) {
	saveItineraryReq := &models.SaveItineraryRequest{}
	if err := c.ShouldBindJSON(saveItineraryReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	saveItineraryResp, err := s.server.SaveItinerary(c, saveItineraryReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, saveItineraryResp)
}

func (s *Service) GetItineraries(c *gin.Context) {
	getItinerariesReq := &models.GetItinerariesRequest{}
	if err := c.ShouldBindJSON(getItinerariesReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	getItinerariesResp, err := s.server.GetItineraries(c, getItinerariesReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getItinerariesResp)
}

func (s *Service) AddReview(c *gin.Context) {
	addReviewReq := &models.AddReviewReq{}
	if err := c.ShouldBindJSON(addReviewReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addReviewResp, err := s.server.AddReview(c, addReviewReq)
	if err != nil {
		if err == server.ErrUserAlreadyCreatedReview {
			c.JSON(http.StatusMethodNotAllowed, err)
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, addReviewResp)
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
