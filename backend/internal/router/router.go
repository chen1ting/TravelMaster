package router

import (
	"github.com/chen1ting/TravelMaster/internal/views"
	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine) {
	r.GET("/ping", views.Ping)
}