package server

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"

	"github.com/google/uuid"

	"github.com/jackc/pgconn"

	gormModel "github.com/chen1ting/TravelMaster/internal/models/gorm"

	"github.com/chen1ting/TravelMaster/internal/models"
)

var (
	ErrBadRequest        = errors.New("bad request")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidLogin      = errors.New("invalid login")
)

func (s *Server) Signup(ctx context.Context, req *models.SignupReq) (*models.SignupResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	file, err := req.Avatar.Open()
	if err!= nil{
		return  nil, err
	}
	defer file.Close()
  
	//new directory "avatars" to save the files
	e := os.MkdirAll("avatars", 0755)
	if e != nil{
		return nil, e
	}
  
	ext:= filepath.Ext(req.Avatar.Filename) //file extension
	n := (uuid.New()).String()              // to random the filename
  
	dest := filepath.Join( "./avatars", n + ext)
  
	///new file locally in avatars
	newfile, err := os.Create(dest)
	if err != nil {
		return nil, err
	}
	defer newfile.Close()
  
	fileBytes, err := io.ReadAll(file)
	if err != nil{
		return nil, err
	}
  
	//copy original file to newfile
	newfile.Write(fileBytes)
  
  
	// attempt to save user to DB
	user := gormModel.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.HashedPassword,
		Avatarurl: dest,
	} 

	if result := s.Database.Model(&user).Create(&user); result.Error != nil {
		// https://github.com/go-gorm/gorm/issues/4135
		var perr *pgconn.PgError
		if errors.As(result.Error, &perr) && perr.Code == "23505" {
			return nil, ErrUserAlreadyExists
		}
		fmt.Println("signup err: ", result.Error) // TODO: write to log instead
		return nil, result.Error
	}

	// create user session
	sessionToken := uuid.New().String()
	if err := s.addNewUserSession(ctx, strconv.Itoa(int(user.ID)), sessionToken, 24*time.Hour); err != nil {
		return nil, err
	}

	return &models.SignupResp{
		UserId:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Avatarurl:    user.Avatarurl,
		SessionToken: sessionToken,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *models.LoginReq) (*models.LoginResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	// query user by username in DB
	var user gormModel.User
	if result := s.Database.Where("Username = ?", req.Username).First(&user); result.Error != nil {
		return nil, ErrInvalidLogin
	}

	// check whether hashed_password matches
	if user.Password != req.HashedPassword {
		return nil, ErrInvalidLogin
	}

	// create user session
	sessionToken := uuid.New().String()
	if err := s.addNewUserSession(ctx, strconv.FormatInt(user.ID, 10), sessionToken, 24*time.Hour); err != nil {
		return nil, err
	}

	return &models.LoginResp{
		UserId:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
		SessionToken: sessionToken,
	}, nil
}

func (s *Server) addNewUserSession(ctx context.Context, userId string, sessionToken string, duration time.Duration) error {
	// delete previous session token if any
	prevToken, err := s.SessionRedis.Get(ctx, userId).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	if err == nil {
		if _, err := s.SessionRedis.Del(ctx, prevToken).Result(); err != nil {
			return err
		}
	}
	if _, err := s.SessionRedis.Set(ctx, userId, sessionToken, duration).Result(); err != nil {
		return err
	}
	if _, err := s.SessionRedis.Set(ctx, sessionToken, userId, duration).Result(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Logout(ctx context.Context, req *models.LogoutReq) error {
	userId, err := s.SessionRedis.Get(ctx, req.SessionToken).Result()
	if err != nil {
		return err
	}

	if _, err := s.SessionRedis.Del(ctx, userId).Result(); err != nil {
		return err
	}
	if _, err := s.SessionRedis.Del(ctx, req.SessionToken).Result(); err != nil {
		return err
	}

	return err
}

func (s *Server) ValidateToken(ctx context.Context, req *models.ValidateTokenReq) (*models.ValidateTokenResp, error) {
	userId, err := s.SessionRedis.Get(ctx, req.SessionToken).Result()
	if err != nil {
		if err == redis.Nil {
			return &models.ValidateTokenResp{
				Valid: false,
				UserId: -1,
			}, nil
		}
		return nil, err
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}
	return &models.ValidateTokenResp{
		Valid: true,
		UserId: uid,
	}, nil
}