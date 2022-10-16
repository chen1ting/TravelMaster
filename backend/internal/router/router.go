package router

import (
	"github.com/chen1ting/TravelMaster/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine, service *service.Service) {
	r.Use(cors.Default()) // default allows all origins

	r.GET("/ping", service.Ping)

	r.POST("/signup", service.SignupView)
	r.POST("/login", service.LoginView)
	r.POST("/get-user-info", service.GetUserInfo) // same function as GetProfile
	r.POST("/get-profile", service.GetProfile)
	r.POST("/update-profile", service.UpdateProfile)
	r.POST("/update-avatar", service.UpdateAvatar)
	r.POST("/logout", service.LogoutView)
	r.POST("/validate-token", service.ValidateToken)
	r.POST("/feedback", service.Feedback)

	r.POST("/generate-itinerary", service.GenerateItinerary)
	r.POST("/get-itinerary", service.GetItinerary)
	r.POST("/save-itinerary", service.UpdateItinerary)
	r.POST("/get-itineraries", service.GetItineraries)

	r.POST("/add-review", service.AddReview)
	r.POST("/update-review", service.UpdateReview)
	r.POST("/create-activity", service.CreateActivity)
	r.POST("/get-activity", service.GetActivity)
	r.POST("/search-activity", service.SearchActivity)
	r.POST("/update-activity", service.UpdateActivity)
	r.DELETE("/delete-activity-image", service.DeleteActivityImage)

	r.POST("/has-user-reported", service.HasUserReported)
	r.POST("/increment-inactive-flag", service.IncrementInactiveCount)
	r.POST("/decrement-inactive-flag", service.DecrementInactiveCount)

	r.Use(static.Serve("/activity-images", static.LocalFile("./assets/activity_images", true)))
	r.Use(static.Serve("/avatars", static.LocalFile("./assets/avatars", true)))
	//r.POST("/create-review", service.CreateReview)

}
