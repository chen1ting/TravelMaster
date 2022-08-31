package main

import (
	"github.com/chen1ting/TravelMaster/internal/router"
	"github.com/chen1ting/TravelMaster/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	s := service.NewService()
	r := gin.Default()
	router.InitRouting(r, s)

	r.Run(":8080")
}
