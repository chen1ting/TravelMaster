package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/jackc/pgconn"

	gormModel "github.com/chen1ting/TravelMaster/internal/models/gorm"

	"github.com/chen1ting/TravelMaster/internal/models"
)

var (
	ErrBadRequest            = errors.New("bad request")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrInvalidLogin          = errors.New("invalid login")
	ErrActivityAlreadyExists = errors.New("an activity with the same title already exists")
	ErrNullTitle             = errors.New("title cannot be empty")
	ErrInvalidActivityID     = errors.New("activity id doesn't exist")
	ErrNoSearchFail          = errors.New("search failed")
	ErrParsingResultFail     = errors.New("cannot parse result")
)

func (s *Server) Signup(ctx context.Context, req *models.SignupReq) (*models.SignupResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	// attempt to save user to DB
	user := gormModel.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.HashedPassword,
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
				Valid:  false,
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
		Valid:  true,
		UserId: uid,
	}, nil
}

func (s *Server) CreateActivity(req *models.CreateActivityReq) (*models.CreateActivityResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	// add activity to database
	activity := gormModel.Activity{
		Title:       req.Title,
		Rating:      req.Rating,
		Paid:        req.Paid,
		Category:    req.Category,
		Description: req.Description,
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,

		MonOpeningTime:  req.MonOpeningTime,
		MonClosingTime:  req.MonClosingTime,
		TueOpeningTime:  req.TueOpeningTime,
		TueClosingTime:  req.TueClosingTime,
		WedOpeningTime:  req.WedOpeningTime,
		WedClosingTime:  req.WedClosingTime,
		ThurOpeningTime: req.ThurOpeningTime,
		ThurClosingTime: req.ThurClosingTime,
		FriOpeningTime:  req.FriOpeningTime,
		FriClosingTime:  req.FriClosingTime,
		SatOpeningTime:  req.SatOpeningTime,
		SatClosingTime:  req.SatClosingTime,
		SunOpeningTime:  req.SunOpeningTime,
		SunClosingTime:  req.SunClosingTime,

		// system settings
		InactiveCount: 0,
		InactiveFlag:  false,
		ReviewCounts:  0,
		ReviewIds:     "",
	}

	if result := s.Database.Model(&activity).Create(&activity); result.Error != nil {
		// error code reference https://github.com/jackc/pgerrcode/blob/master/errcode.go
		var perr *pgconn.PgError
		if errors.As(result.Error, &perr) && perr.Code == "23505" {
			return nil, ErrActivityAlreadyExists
		}
		if errors.As(result.Error, &perr) && perr.Code == "23502" {
			return nil, ErrNullTitle
		}
		fmt.Println("create_activity err: ", result.Error) // TODO: write to log instead
		return nil, result.Error
	}

	// if no error, return success response
	return &models.CreateActivityResp{
		ActivityId: activity.ID,
		CreatedAt:  activity.CreatedAt,
	}, nil
}

func (s *Server) GetActivity(req *models.GetActivityReq) (*models.GetActivityResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}
	var activity gormModel.Activity

	// if activity cannot be found by given ID, return error
	if result := s.Database.First(&activity, req.ActivityId); result.Error != nil {
		return nil, ErrInvalidActivityID
	}

	return &models.GetActivityResp{
		ActivityId:  activity.ID,
		Title:       activity.Title,
		Rating:      activity.Rating,
		Paid:        activity.Paid,
		Category:    activity.Category,
		Description: activity.Description,
		Longitude:   activity.Longitude,
		Latitude:    activity.Latitude,
		ImageURL:    activity.ImageURL,

		MonOpeningTime:  activity.MonOpeningTime,
		MonClosingTime:  activity.MonClosingTime,
		TueOpeningTime:  activity.TueOpeningTime,
		TueClosingTime:  activity.TueClosingTime,
		WedOpeningTime:  activity.WedOpeningTime,
		WedClosingTime:  activity.WedClosingTime,
		ThurOpeningTime: activity.ThurOpeningTime,
		ThurClosingTime: activity.ThurClosingTime,
		FriOpeningTime:  activity.FriOpeningTime,
		FriClosingTime:  activity.FriClosingTime,
		SatOpeningTime:  activity.SatOpeningTime,
		SatClosingTime:  activity.SatClosingTime,
		SunOpeningTime:  activity.SunOpeningTime,
		SunClosingTime:  activity.SunClosingTime,

		InactiveCount: activity.InactiveCount,
		InactiveFlag:  activity.InactiveFlag,
		ReviewCounts:  activity.ReviewCounts,
		//ReviewList: activity.ReviewID,	// TODO: get reviews
		CreatedAt: activity.CreatedAt,
	}, nil
}

func Paginate(r *models.SearchActivityReq) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := r.PageNumber
		if page <= 0 {
			page = 1
		}

		pageSize := r.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *Server) SearchActivity(req *models.SearchActivityReq) (*models.SearchActivityResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}
	var activities []gormModel.Activity

	result := s.Database.Where("title ILIKE ? AND inactive_flag = ?", "%"+req.SearchText+"%", "0").Order("created_at desc").Scopes(Paginate(req)).Find(&activities)
	// if activity cannot be found by given ID, return error
	if result.Error != nil {
		return nil, ErrNoSearchFail
	}
	jsStr, err := json.Marshal(activities)
	if err != nil {
		return nil, ErrParsingResultFail
	}
	return &models.SearchActivityResp{
		Activities:   string(jsStr),
		ResultNumber: result.RowsAffected,
	}, nil
}

func (s *Server) UpdateActivity(req *models.UpdateActivityReq) (*models.UpdateActivityResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	if req.Title == "" {
		return nil, ErrNullTitle
	}

	// find the activity in database
	var activity gormModel.Activity

	// if activity cannot be found by given ID, return error
	if result := s.Database.First(&activity, req.ActivityId); result.Error != nil {
		return nil, ErrInvalidActivityID
	}

	// update activity and save to database
	activity.Title = req.Title
	activity.Rating = req.Rating
	activity.Paid = req.Paid
	activity.Category = req.Category
	activity.Description = req.Description
	activity.Longitude = req.Longitude
	activity.Latitude = req.Latitude
	// TODO: update image

	activity.MonOpeningTime = req.MonOpeningTime
	activity.MonClosingTime = req.MonClosingTime
	activity.TueOpeningTime = req.TueOpeningTime
	activity.TueClosingTime = req.TueClosingTime
	activity.WedOpeningTime = req.WedOpeningTime
	activity.WedClosingTime = req.WedClosingTime
	activity.ThurOpeningTime = req.ThurOpeningTime
	activity.ThurClosingTime = req.ThurClosingTime
	activity.FriOpeningTime = req.FriOpeningTime
	activity.FriClosingTime = req.FriClosingTime
	activity.SatOpeningTime = req.SatOpeningTime
	activity.SatClosingTime = req.SatClosingTime
	activity.SunOpeningTime = req.SunOpeningTime
	activity.SunClosingTime = req.SunClosingTime

	s.Database.Save(&activity)

	if result := s.Database.Save(&activity); result.Error != nil {
		var perr *pgconn.PgError
		if errors.As(result.Error, &perr) && perr.Code == "23505" {
			return nil, ErrActivityAlreadyExists
		}
		fmt.Println("update_activity err: ", result.Error) // TODO: write to log instead
		return nil, result.Error
	}

	// if no error, return success response
	return &models.UpdateActivityResp{
		ActivityId: activity.ID,
		UpdatedAt:  activity.UpdatedAt,
	}, nil
}

func (s *Server) ReportInactiveActivity(req *models.InactivateActivityReq) (*models.InactivateActivityResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	// find the activity in database
	var activity gormModel.Activity

	// if activity cannot be found by given ID, return error
	if result := s.Database.First(&activity, req.ActivityId); result.Error != nil {
		return nil, ErrInvalidActivityID
	}

	//TODO: put invalid threshold into global variable?
	invalidThreshold := 10

	activity.InactiveCount++
	if activity.InactiveCount >= invalidThreshold {
		activity.InactiveFlag = true
	}

	if result := s.Database.Save(&activity); result.Error != nil {
		fmt.Println("inactivate_activity err: ", result.Error) // TODO: write to log instead
		return nil, result.Error
	}

	return &models.InactivateActivityResp{
		ActivityId:    activity.ID,
		InactiveCount: activity.InactiveCount,
		InactiveFlag:  activity.InactiveFlag,
		UpdatedAt:     activity.UpdatedAt,
	}, nil
}
