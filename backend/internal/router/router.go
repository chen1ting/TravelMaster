package router

import (
	"github.com/chen1ting/TravelMaster/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine, service *service.Service) {
	r.Use(cors.Default()) // default allows all origins

	r.GET("/ping", service.Ping)
	r.POST("/signup", service.SignupView)
	r.POST("/login", service.LoginView)
	r.POST("/logout", service.LogoutView)
}
