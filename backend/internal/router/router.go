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
	r.POST("/logout", service.LogoutView)
	r.POST("/validate-token", service.ValidateToken)
	r.POST("/generate-itinerary", service.GenerateItinerary)
	r.POST("/get-itinerary", service.GetItinerary)
	r.POST("/save-itinerary", service.UpdateItinerary)
	r.POST("/get-itineraries", service.GetItineraries)
	r.POST("/create-activity", service.CreateActivity)
	r.POST("/add-review", service.AddReview)
	r.POST("/get-activity", service.GetActivity)
	r.POST("/search-activity", service.SearchActivity)
	r.POST("/update-activity", service.UpdateActivity)
	r.POST("/report-inactive-activity", service.ReportInactiveActivity)
	r.Use(static.Serve("/activity-images", static.LocalFile("./assets/activity_images", true)))
	r.Use(static.Serve("/avatars", static.LocalFile("./assets/avatars", true)))
	r.DELETE("/delete-activity-image", service.DeleteActivityImage)
}
