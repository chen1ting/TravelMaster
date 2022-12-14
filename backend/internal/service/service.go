package service

import (
	"fmt"
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

func NewService(env string) *Service {
	return &Service{
		server: server.NewServer(env),
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
	fmt.Println(signupResp)

	c.JSON(http.StatusCreated, signupResp)
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
		if err == server.ErrBadRequest {
			c.JSON(http.StatusBadRequest, err)
			return
		}
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

func (s *Service) UpdateProfile(c *gin.Context) {
	updateProfileReq := &models.UpdateProfileReq{}
	if err := c.ShouldBindJSON(&updateProfileReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createActivityResp, err := s.server.UpdateProfile(updateProfileReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, createActivityResp)
}

func (s *Service) GetProfile(c *gin.Context) {
	getProfileReq := &models.GetProfileReq{}
	if err := c.ShouldBindJSON(&getProfileReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createActivityResp, err := s.server.GetProfile(getProfileReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createActivityResp)
}

func (s *Service) UpdateAvatar(c *gin.Context) {
	updateAvatarForm := &models.UpdateAvatarForm{}
	if err := c.ShouldBind(&updateAvatarForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createActivityResp, err := s.server.UpdateAvatar(updateAvatarForm, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, createActivityResp)
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
			c.JSON(http.StatusMethodNotAllowed, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, addReviewResp)
}

func (s *Service) GetUserInfo(c *gin.Context) {
	req := &models.GetUserInfoReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := s.server.GetUserInfo(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
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

func (s *Service) IncrementInactiveCount(c *gin.Context) {
	incrementInactiveCountReq := &models.IncrementInactiveCountReq{}
	if err := c.ShouldBindJSON(&incrementInactiveCountReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inactivateActivityResp, err := s.server.IncrementInactiveCount(incrementInactiveCountReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, inactivateActivityResp)
}

func (s *Service) DecrementInactiveCount(c *gin.Context) {
	decrementInactiveCountReq := &models.DecrementInactiveCountReq{}
	if err := c.ShouldBindJSON(&decrementInactiveCountReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inactivateActivityResp, err := s.server.DecrementInactiveCount(decrementInactiveCountReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, inactivateActivityResp)
}

func (s *Service) HasUserReported(c *gin.Context) {
	hasUserInactivatedReq := &models.HasUserInactivatedReq{}
	if err := c.ShouldBindJSON(&hasUserInactivatedReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hasUserInactivatedResp, err := s.server.CheckInactiveFlag(hasUserInactivatedReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, hasUserInactivatedResp)
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

/*
func (s *Service) CreateReview(c *gin.Context) {
	createReviewReq := &models.CreateReviewReq{}
	if err := c.ShouldBindJSON(&createReviewReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createReviewResp, err := s.server.CreateReview(createReviewReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, createReviewResp)
}*/

func (s *Service) UpdateReview(c *gin.Context) {
	updateReviewReq := &models.UpdateReviewReq{}
	if err := c.ShouldBindJSON(&updateReviewReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	getActivityResp, err := s.server.UpdateReview(updateReviewReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, getActivityResp)
}

// blank endpoint
func (s *Service) Feedback(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "ok"})
}
