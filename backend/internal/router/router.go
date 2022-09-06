package router

import (
	"github.com/chen1ting/TravelMaster/internal/service"
	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine, service *service.Service) {
	r.GET("/ping", service.Ping)
	r.POST("/signup", service.SignupView)
}
