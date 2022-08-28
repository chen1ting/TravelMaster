package main

import (
	"github.com/chen1ting/TravelMaster/internal/config"
	"github.com/chen1ting/TravelMaster/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	r := gin.Default()
	router.InitRouting(r)

	r.Run(":8080")
}