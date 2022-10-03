package server

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"

	"github.com/chen1ting/TravelMaster/internal/config"
	"github.com/chen1ting/TravelMaster/internal/models"
	gormModel "github.com/chen1ting/TravelMaster/internal/models/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	Config       config.Config
	Database     *gorm.DB
	SessionRedis *redis.Client
}

type ServerInf interface {
	Signup(c *gin.Context, req *models.SignupForm) (*models.SignupResp, error)
	Login(ctx context.Context, req *models.LoginReq) (*models.LoginResp, error)
	// TODO: add in all handle functions here?
}

var _ ServerInf = (*Server)(nil)

func NewServer() *Server {
	conf := config.NewConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", conf.DBHost, conf.DBUser, conf.DBPass, conf.DBName, conf.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&gormModel.User{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&gormModel.Activity{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&gormModel.Review{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&gormModel.Itinerary{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&gormModel.Review{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&gormModel.ReportHistory{}); err != nil {
		panic(err)
	}
	if err := db.SetupJoinTable(&gormModel.Activity{}, "UserReports", &gormModel.ReportHistory{}); err != nil {
		panic(err)
	}

	sessionRedis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.SessionRedisHost, conf.SessionRedisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Server{
		Config:       conf,
		Database:     db,
		SessionRedis: sessionRedis,
	}
}
