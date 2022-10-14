package main

import (
	"os"

	"github.com/chen1ting/TravelMaster/internal/router"
	"github.com/chen1ting/TravelMaster/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	s := service.NewService()
	r := gin.Default()
	router.InitRouting(r, s)

	if os.Getenv("APP_ENV") == "development" {
		r.Run("0.0.0.0:8080")
	} else {
		r.RunTLS("0.0.0.0:8080", "server.pem", "key.unencrypted.pem")
	}
}
