package main

import (
	"os"

	"github.com/chen1ting/TravelMaster/internal/router"
	"github.com/chen1ting/TravelMaster/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	s := service.NewService(os.Getenv("APP_ENV"))
	router.InitRouting(r, s)

	if os.Getenv("APP_ENV") == "development" || os.Getenv("APP_ENV") == "testing" {
		r.Run("0.0.0.0:8080")
	} else {
		r.RunTLS("0.0.0.0:8080", "server.pem", "key.unencrypted.pem")
	}
}
