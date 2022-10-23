package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/google/uuid"

	gormModel "github.com/chen1ting/TravelMaster/internal/models/gorm"

	"github.com/chen1ting/TravelMaster/internal/models"
)

var (
	ErrUserNotExist             = errors.New("user not exist")
	ErrUserAlreadyCreatedReview = errors.New("user already created review for the activity")
	ErrNotAllowed               = errors.New("user is not allowed to perform this action")
	ErrGenericServerError       = errors.New("generic server error")
	ErrDatabase                 = errors.New("database error")
	ErrBadRequest               = errors.New("bad request")
	ErrMissingUserInfo          = errors.New("either email, username, or hashed password missing")
	ErrUserAlreadyExists        = errors.New("user already exists")
	ErrInvalidLogin             = errors.New("invalid login")
	ErrActivityAlreadyExists    = errors.New("an activity with the same title already exists")
	ErrActivityNotFound         = errors.New("activity not found")
	ErrNullTitle                = errors.New("title cannot be empty")
	ErrUserNotFound             = errors.New("user id doesn't exists")
	ErrReviewNotFound           = errors.New("review not found")
	ErrReportNotFound           = errors.New("report not found")
	ErrInvalidUpdateUser        = errors.New("user id doesn't match the activity's user id")
	ErrNoSearchFail             = errors.New("search Name failed")
	ErrAlreadyReported          = errors.New("user has already reported the activity")
	ErrUnknownFileType          = errors.New("unknown file type uploaded")
	ErrImageNoMatch             = errors.New("image not found in the list of the activity")
	ErrImageNotFound            = errors.New("image not found on server, removed file name in the database")
	CWD, _                      = os.Getwd()
	ImageRoot                   = filepath.Join(CWD, "assets")
	ActivityImageFolder         = "activity_images"
	AvatarFolder                = "avatars"
	InvalidThreshold            = 10
)

func (s *Server) Signup(c *gin.Context, form *models.SignupForm) (*models.SignupResp, error) {
	if form == nil {
		return nil, ErrBadRequest
	}

	// assumptions: email, username, and password cannot be empty
	if form.Email == "" || form.Username == "" || form.HashedPassword == "" {
		return nil, ErrMissingUserInfo
	}

	// validate user doesn't exist before saving images
	var user gormModel.User
	if result := s.Database.Where("username = ? OR email = ?", form.Username, form.Email).First(&user); result.RowsAffected != 0 {
		return nil, ErrUserAlreadyExists
	}

	//check  if there is file first
	var uniqueImgName, fpath string
	if form.Avatar != nil {
		imgName, filePath, saveErr := SaveFile(form.Avatar, c, AvatarFolder)
		if saveErr != nil {
			fmt.Println(saveErr) //TODO: log instead
		}
		uniqueImgName = imgName
		fpath = filePath
	}

	// attempt to save user to DB
	user = gormModel.User{
		Username:   form.Username,
		Email:      form.Email,
		Password:   form.HashedPassword,
		AvatarName: uniqueImgName,
	}

	if result := s.Database.Omit(clause.Associations).Create(&user); result.Error != nil { //s.Database.Model(&user).Create(&user)
		err := os.Remove(fpath)
		if err != nil { // TODO: write to log instead
			if err := os.Remove(fpath); err != nil {
				fmt.Println("sign_up have error deleting avatar: ", err)
				return nil, err
			} // delete image
		}
		return nil, result.Error
	}

	// create user session
	sessionToken := uuid.New().String()
	if err := s.addNewUserSession(c, strconv.Itoa(int(user.ID)), sessionToken, 24*time.Hour); err != nil {
		return nil, err
	}
	return &models.SignupResp{
		UserId:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
		AvatarName:   user.AvatarName,
		SessionToken: sessionToken,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *models.LoginReq) (*models.LoginResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	// query user by username in DB
	var user gormModel.User
	if result := s.Database.Where("username = ?", req.Username).First(&user); result.Error != nil {
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
		AvatarName:   user.AvatarName,
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

func (s *Server) GenerateItinerary(ctx context.Context, req *models.GenerateItineraryRequest) (*models.GenerateItineraryResponse, error) {
	if req.StartTime > req.EndTime {
		return nil, ErrBadRequest
	}
	userId, err := s.SessionRedis.Get(ctx, req.SessionToken).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotAllowed
		}
		return nil, err
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	// retrieve all activities
	var activities []gormModel.Activity
	if err := s.Database.Find(&activities).Error; err != nil {
		return nil, ErrDatabase
	}

	actMap := make(map[string][]*gormModel.Activity)
	for _, act := range activities {
		for _, cat := range act.Categories {
			if actMap[cat] == nil {
				actMap[cat] = make([]*gormModel.Activity, 0)
			}
			actMap[cat] = append(actMap[cat], &gormModel.Activity{
				ID:            act.ID,
				UserID:        act.UserID,
				Title:         act.Title,
				AverageRating: act.AverageRating,
				Paid:          act.Paid,
				Categories:    act.Categories,
				Description:   act.Description,
				Longitude:     act.Longitude,
				Latitude:      act.Latitude,
				ImageNames:    act.ImageNames,
				OpeningTimes:  act.OpeningTimes,
				InactiveCount: act.InactiveCount,
				InactiveFlag:  act.InactiveFlag,
				ReviewCounts:  act.ReviewCounts,
				Reviews:       act.Reviews,
				CreatedAt:     act.CreatedAt,
				UpdatedAt:     act.UpdatedAt,
			})
		}
	}

	// fill up fixed slots first 8am - 10am
	startTime := time.Unix(req.StartTime, 0)
	startBase := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location()).Unix()
	// endTime := time.Unix(req.EndTime, 0)
	// endBase := time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, endTime.Location()).Unix()
	// n := (endBase - startBase) / (60*60) + 24
	// buckets := make([]*models.Segment, n)
	startIdx := (req.StartTime - startBase) / (60 * 60)
	day := int(float64(req.StartTime/86400)+4) % 7
	// endIdx := (req.EndTime - startBase) / (60*60)
	hr := int(startIdx) // 0 indexed hr
	x := req.StartTime
	y := req.EndTime
	segments := make([]*models.Segment, 0)
	used := make(map[int64]bool)
	for x <= y {
		if (hr >= 7 && hr <= 8) || (hr >= 11 && hr <= 12) || (hr >= 6 && hr <= 7) { // breakfast, lunch, time
			// randomly select a food activity that is open at that time
			activity, h := randomAndIsOpen(actMap["Food and beverage"], day, hr, used)
			if activity == nil { // no food activity somehow...
				fmt.Printf("WARN: no food activity for start time: %d\n", x)
				hr += 1
				x += int64(60 * 60)
				continue
			}
			segments = append(segments, &models.Segment{
				StartTime:       x,
				EndTime:         x + int64(h*60*60),
				ActivitySummary: activity,
			})
			hr += h + 2 // 2h gap between every activity
			x += int64((h + 2) * 60 * 60)
		} else if hr >= 21 { // 10 PM or later, fast-forward to 8 AM next day
			ff := 7 + 24 - hr
			hr = 7
			x += int64(ff * 60 * 60)
			day = (day + 1) % 7
		} else { // any 2 hr time slot for any activity or less if exceeds end time
			ok := false
			for _, cat := range req.PreferredCategories {
				activity, h := randomAndIsOpen(actMap[cat], day, hr, used)
				if activity != nil {
					ok = true
					segments = append(segments, &models.Segment{
						StartTime:       x,
						EndTime:         x + int64(h*60*60),
						ActivitySummary: activity,
					})
					hr += h + 2 // 2h gap between every activity
					x += int64((h + 2) * 60 * 60)
					break
				}
			}
			if !ok {
				// any cat is ok
				for _, m := range actMap {
					activity, h := randomAndIsOpen(m, day, hr, used)
					if activity != nil {
						ok = true
						segments = append(segments, &models.Segment{
							StartTime:       x,
							EndTime:         x + int64(h*60*60),
							ActivitySummary: activity,
						})
						hr += h + 2 // 2h gap between every activity
						x += int64((h + 2) * 60 * 60)
						break
					}
				}
			}
			if !ok {
				fmt.Printf("WARN: no planned activity for start time: %d\n", x)
				hr += 1
				x += int64(60 * 60)
			}
		}
	}

	marshalledSeg, err := json.Marshal(segments)
	if err != nil {
		return nil, ErrGenericServerError
	}
	// insert into db generated itinerary
	genIt := &gormModel.Itinerary{
		Name:             uuid.New().String(),
		OwnedByUserId:    uid,
		Segments:         marshalledSeg,
		StartTime:        req.StartTime,
		EndTime:          req.EndTime,
		NumberOfSegments: len(segments),
	}
	if res := s.Database.Create(genIt); res.Error != nil {
		return nil, ErrDatabase
	}

	// return itinerary as resp
	return &models.GenerateItineraryResponse{
		GeneratedItinerary: &models.Itinerary{
			Id:               genIt.ID,
			Name:             genIt.Name,
			NumberOfSegments: len(segments),
			Segments:         segments,
			StartTime:        req.StartTime,
			EndTime:          req.EndTime,
		},
	}, nil
}

func (s *Server) SaveItinerary(ctx context.Context, req *models.SaveItineraryRequest) (resp *models.SaveItineraryResponse, err error) {
	userId, err := s.SessionRedis.Get(ctx, req.SessionToken).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotAllowed
		}
		return nil, err
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}
	var iti gormModel.Itinerary
	if res := s.Database.Find(&iti, req.Id); res.Error != nil {
		return nil, ErrDatabase
	}
	if iti.OwnedByUserId != uid {
		return nil, ErrNotAllowed
	}

	marshalledSeg, err := json.Marshal(req.Segments)
	if err != nil {
		return nil, ErrGenericServerError
	}
	iti.Name = req.Name
	iti.Segments = marshalledSeg
	if res := s.Database.Save(&iti); res.Error != nil {
		return nil, ErrDatabase
	}

	return &models.SaveItineraryResponse{Id: iti.ID, Name: iti.Name}, nil
}

// returns the activity summary and the time allocated for the activity: 1 or 2 hr
// it will try to return 2h, and only return 1 if the activity ends before x+2
func randomAndIsOpen(choices []*gormModel.Activity, day int, hr int, used map[int64]bool) (*models.ActivitySummary, int) {
	for _, act := range choices {
		opening := int(act.OpeningTimes[day])
		closing := int(act.OpeningTimes[day+7])
		if hr < opening || hr > closing || used[act.ID] {
			continue
		}
		actTime := min(2, closing-hr)
		if actTime == 0 { // act time must at least an hour long
			continue
		}
		imageUrl := ""
		if len(act.ImageNames) > 0 {
			imageUrl = act.ImageNames[0]
		}
		used[act.ID] = true
		return &models.ActivitySummary{
			Id:            act.ID,
			Name:          act.Title,
			Description:   act.Description,
			AverageRating: Round(float64(act.AverageRating), 0.05),
			Categories:    act.Categories,
			ImageNames:    []string{imageUrl},
			ReviewCounts:  act.ReviewCounts,
		}, actTime
	}

	return nil, 0
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func (s *Server) GetItinerary(ctx context.Context, req *models.GetItineraryRequest) (*models.GetItineraryResponse, error) {
	userId, err := s.SessionRedis.Get(ctx, req.SessionToken).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotAllowed
		}
		return nil, err
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}
	var iti gormModel.Itinerary
	if res := s.Database.Find(&iti, req.Id); res.Error != nil {
		return nil, ErrDatabase
	}
	if iti.OwnedByUserId != uid {
		return nil, ErrNotAllowed
	}

	var segments []*models.Segment
	if err := json.Unmarshal(iti.Segments, &segments); err != nil {
		fmt.Println(err)
		return nil, ErrGenericServerError
	}

	return &models.GetItineraryResponse{
		Itinerary: &models.Itinerary{
			Id:               iti.ID,
			Name:             iti.Name,
			NumberOfSegments: iti.NumberOfSegments,
			Segments:         segments,
			StartTime:        iti.StartTime,
			EndTime:          iti.EndTime,
		},
	}, nil
}

func (s *Server) GetItineraries(ctx context.Context, req *models.GetItinerariesRequest) (*models.GetItinerariesResponse, error) {
	userId, err := s.SessionRedis.Get(ctx, req.SessionToken).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotAllowed
		}
		return nil, err
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}
	var itis []gormModel.Itinerary
	if res := s.Database.Where("owned_by_user_id = ?", uid).Find(&itis); res.Error != nil {
		return nil, ErrDatabase
	}

	parsedItis := make([]*models.Itinerary, 0)
	for _, iti := range itis {
		var segments []*models.Segment
		if err := json.Unmarshal(iti.Segments, &segments); err != nil {
			fmt.Println(err)
			return nil, ErrGenericServerError
		}
		parsedItis = append(parsedItis, &models.Itinerary{
			Id:               iti.ID,
			Name:             iti.Name,
			NumberOfSegments: iti.NumberOfSegments,
			Segments:         segments,
			StartTime:        iti.StartTime,
			EndTime:          iti.EndTime,
		})
	}

	return &models.GetItinerariesResponse{Itineraries: parsedItis}, nil
}

func getValidTime(hhmm int) int {
	hour := hhmm / 100
	min := hhmm % 100
	if 0 <= hour && hour < 24 && 0 <= min && min < 60 {
		return hhmm
	}
	return -1
}

func (s *Server) UpdateProfile(req *models.UpdateProfileReq) (*models.UpdateProfileResp, error) {
	var user gormModel.User
	if result := s.Database.First(&user, req.UserId); result.Error != nil {
		return nil, ErrUserNotFound
	}
	user.Interests = req.Interests
	user.AboutMe = req.AboutMe

	if result := s.Database.Save(&user); result.Error != nil {
		fmt.Println("update_profile err: ", result.Error) // TODO: write to log instead
		return nil, result.Error
	}
	return &models.UpdateProfileResp{
		UserId:    user.ID,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *Server) GetProfile(req *models.GetProfileReq) (*models.GetProfileResp, error) {
	var user gormModel.User
	if result := s.Database.Where("id=?", req.UserId).Preload("Activities").Preload("Reviews").Find(&user); result.RowsAffected == 0 {
		//if result := s.Database.First(&user, req.UserId); result.Error != nil {
		return nil, ErrUserNotFound
	}

	// parse user activities into response
	parsedActivities := make([]*models.GetActivityResp, 0)
	for _, activity := range user.Activities {
		parsedActivities = append(parsedActivities, ParseActivity(activity))
	}

	return &models.GetProfileResp{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		AboutMe:     user.AboutMe,
		AvatarName:  user.AvatarName,
		Activities:  parsedActivities,
		Reviews:     ParseReviewList(user.Reviews),
		CreatedAt:   user.CreatedAt,
		RetrievedAt: time.Now(),
	}, nil
}

func (s *Server) UpdateAvatar(form *models.UpdateAvatarForm, c *gin.Context) (*models.UpdateAvatarResp, error) {
	var user gormModel.User
	if result := s.Database.First(&user, form.UserId); result.Error != nil {
		return nil, ErrUserNotFound
	}
	// assumption: delete request has higher priority
	if form.Delete {
		if err := os.Remove(filepath.Join(ImageRoot, AvatarFolder, user.AvatarName)); err != nil {
			fmt.Println("Unable to remove user avatar, ", err.Error())
		}
		user.AvatarName = "" //reset the avatar name
		if result := s.Database.Save(&user); result.Error != nil {
			fmt.Println("Unable to update database", result.Error.Error())
			return nil, result.Error
		}
		return &models.UpdateAvatarResp{
			UserId:            user.ID,
			UpdatedAt:         time.Now(),
			NewAvatarFileName: "",
		}, nil
	}
	// if the request is not delete, yet no avatar file received, return error
	if form.Avatar == nil {
		fmt.Println("received no file for update", ErrMissingUserInfo.Error())
		return nil, ErrMissingUserInfo
	}
	// if saveErr occurs, return nil
	var avatarName, avatarPath string
	if form.Avatar != nil {
		avaName, avaPath, saveErr := SaveFile(form.Avatar, c, AvatarFolder)
		if saveErr != nil {
			return nil, saveErr
		}
		avatarName = avaName
		avatarPath = avaPath
	}
	user.AvatarName = avatarName
	if result := s.Database.Save(&user); result.Error != nil {
		if err := os.Remove(avatarPath); err != nil {
			fmt.Println("received no file for update", ErrMissingUserInfo.Error())
			return nil, err
		}
		return nil, result.Error
	}
	return &models.UpdateAvatarResp{
		UserId:            user.ID,
		UpdatedAt:         user.UpdatedAt,
		NewAvatarFileName: user.AvatarName,
	}, nil
}

func (s *Server) CreateActivity(form *models.CreateActivityForm, c *gin.Context) (*models.CreateActivityResp, error) {
	if form == nil {
		return nil, ErrBadRequest
	}

	var user gormModel.User
	if result := s.Database.First(&user, form.UserId); result.Error != nil {
		return nil, ErrUserNotFound
	}

	if form.Title == "" {
		fmt.Println("create_activity err: ", ErrNullTitle) // TODO: write to log instead
		return nil, ErrNullTitle
	}

	var activitySearch gormModel.Activity
	if result := s.Database.Where("Title = ?", form.Title).First(&activitySearch); result.RowsAffected > 0 {
		fmt.Println("create_activity err: ", ErrActivityAlreadyExists) // TODO: write to log instead
		return nil, ErrActivityAlreadyExists
	}

	var imgNames []string
	var imgPaths []string
	var failedImages []string
	if form.Image != nil {
		for i := 0; i < len(form.Image); i++ {
			uniqueImgName, fpath, saveErr := SaveFile(form.Image[i], c, ActivityImageFolder)
			if saveErr != nil {
				failedImages = append(failedImages, form.Image[i].Filename)
				continue
			}
			imgNames = append(imgNames, uniqueImgName)
			imgPaths = append(imgPaths, fpath)
		}
	}

	// add activity to database
	activity := gormModel.Activity{
		UserID:       form.UserId,
		Title:        form.Title,
		Paid:         form.Paid,
		AuthorRating: form.Rating,
		Categories:   form.Categories,
		Description:  form.Description,
		Longitude:    form.Longitude,
		Latitude:     form.Latitude,
		OpeningTimes: PackCreateOpeningTimes(form),
		ImageNames:   imgNames,
	}
	//if err := s.Database.Model(&user).Association("Activities").Append(&activity); err != nil {
	if result := s.Database.Save(&activity); result.Error != nil {
		fmt.Println("create_activity err: ", result.Error) // TODO: write to log instead
		// if result cannot be saved, RemoveName all saved images
		for i := 0; i < len(imgPaths); i++ {
			err := os.Remove(imgPaths[i])
			if err != nil { // TODO: write to log instead
				fmt.Println("create_activity error deleting image file: ", err)
			}
		}
		return nil, result.Error
	}

	// if no error, return success response
	return &models.CreateActivityResp{
		ActivityId:     activity.ID,
		CreatedAt:      activity.CreatedAt,
		ImageSaveFails: failedImages,
	}, nil
}

func (s *Server) UpdateActivity(form *models.UpdateActivityForm, c *gin.Context) (*models.UpdateActivityResp, error) {
	if form == nil {
		return nil, ErrBadRequest
	}

	// find the activity in database
	var activity gormModel.Activity

	// if activity cannot be found by given ID, return error
	if result := s.Database.First(&activity, form.ActivityId); result.Error != nil || result.RowsAffected == 0 {
		return nil, ErrActivityNotFound
	}

	// see if the user id matches the activity's user id
	if activity.UserID != form.UserId {
		return nil, ErrInvalidUpdateUser
	}

	if form.Title == "" {
		return nil, ErrNullTitle
	}

	var imgNames = activity.ImageNames
	var imgPaths []string
	var failedImages []string
	if form.Image != nil {
		for i := 0; i < len(form.Image); i++ {
			uniqueImgName, fpath, saveErr := SaveFile(form.Image[i], c, ActivityImageFolder)
			if saveErr != nil {
				failedImages = append(failedImages, form.Image[i].Filename)
				continue
			}
			imgNames = append(imgNames, uniqueImgName)
			imgPaths = append(imgPaths, fpath)
		}
	}

	// update activity and save to database
	activity.Title = form.Title
	activity.AuthorRating = form.Rating
	activity.Paid = form.Paid
	activity.Categories = form.Categories
	activity.Description = form.Description
	activity.Longitude = form.Longitude
	activity.Latitude = form.Latitude
	activity.ImageNames = imgNames
	activity.OpeningTimes = PackUpdateOpeningTimes(form)

	if result := s.Database.Save(&activity); result.Error != nil {
		for i := 0; i < len(imgPaths); i++ {
			err := os.Remove(imgPaths[i])
			if err != nil { // TODO: write to log instead
				fmt.Println("create_activity error deleting image file: ", err)
			}
		}
		fmt.Println("update_activity err: ", result.Error) // TODO: write to log instead
		return nil, result.Error
	}
	// if no error, return success response
	return &models.UpdateActivityResp{
		ActivityId:     activity.ID,
		UpdatedAt:      activity.UpdatedAt,
		ImageSaveFails: failedImages,
	}, nil
}

func (s *Server) GetUserInfo(ctx context.Context, req *models.GetUserInfoReq) (*models.GetUserInfoResp, error) {
	var user gormModel.User
	if res := s.Database.First(&user, req.UserId); res.Error != nil {
		// TODO: for now assume user not exist, and not some db conn/query issue
		return nil, ErrUserNotExist
	}

	return &models.GetUserInfoResp{Username: user.Username, AvatarUrl: user.AvatarName}, nil
}

// TODO: IMPT: this function is not safe for concurrent access, we should implement a lock
func (s *Server) AddReview(ctx context.Context, req *models.AddReviewReq) (*models.GetActivityResp, error) {
	userId, err := s.SessionRedis.Get(ctx, req.SessionToken).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotAllowed
		}
		return nil, err
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	// fetch activity
	var activity gormModel.Activity
	result := s.Database.Where("id=?", req.ActivityId).Preload("Reviews").Find(&activity) //s.Database.First(&activity, req.ActivityId)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, ErrActivityNotFound
	}

	// assert that user has not made a review before for this activity
	review := gormModel.Review{
		Title:       req.Title,
		Description: req.Description,
		UserId:      uid,
		ActivityId:  req.ActivityId,
		Rating:      req.Rating,
	}
	if res := s.Database.Create(&review); res.Error != nil {
		// i'll just assume its a violation error here, but its not necessarily the case
		return nil, ErrUserAlreadyCreatedReview
	}

	newAvg := (float32(activity.ReviewCounts)*activity.AverageRating + req.Rating) / float32(activity.ReviewCounts+1)
	activity.ReviewCounts++
	activity.AverageRating = newAvg
	if res := s.Database.Save(&activity); res.Error != nil {
		return nil, ErrDatabase
	}

	return ParseActivity(activity), nil
}

func (s *Server) GetActivity(req *models.GetActivityReq) (*models.GetActivityResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}
	var activity gormModel.Activity

	// if activity cannot be found by given ID, return error
	if result := s.Database.Where("id=?", req.ActivityId).Preload("Reviews").Find(&activity); result.Error != nil || result.RowsAffected == 0 {
		return nil, ErrActivityNotFound
	}

	return ParseActivity(activity), nil
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

	// one more level of filtering by times allowed
	// only enforced if specified
	filteredAct := make([]*models.ActivitySummary, 0)
	if len(req.Times) > 0 {
		for _, act := range activities {
			ok := true
			for _, time := range req.Times {
				if !(time.StartTimeOffset >= int(act.OpeningTimes[time.Day]) && time.EndTimeOffset <= int(act.OpeningTimes[time.Day+7])) {
					fmt.Println("not ok: ", time.StartTimeOffset, int(act.OpeningTimes[time.Day]), time.EndTimeOffset, int(act.OpeningTimes[time.Day+7]))
					ok = false
					break
				}
			}
			imageUrl := ""
			if len(act.ImageNames) > 0 {
				imageUrl = act.ImageNames[0]
			}
			if ok {
				filteredAct = append(filteredAct, &models.ActivitySummary{
					Id:            act.ID,
					Name:          act.Title,
					Description:   act.Description,
					AverageRating: Round(float64(act.AverageRating), 0.05),
					Categories:    act.Categories,
					ImageNames:    []string{imageUrl},
					ReviewCounts:  act.ReviewCounts,
				})
			}
		}
	} else {
		for _, act := range activities {
			imageUrl := ""
			if len(act.ImageNames) > 0 {
				imageUrl = act.ImageNames[0]
			}
			filteredAct = append(filteredAct, &models.ActivitySummary{
				Id:            act.ID,
				Name:          act.Title,
				Description:   act.Description,
				AverageRating: Round(float64(act.AverageRating), 0.05),
				Categories:    act.Categories,
				ImageNames:    []string{imageUrl},
				ReviewCounts:  act.ReviewCounts,
			})
		}
	}

	return &models.SearchActivityResp{
		Activities:   filteredAct,
		NumOfResults: len(filteredAct),
	}, nil
}

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func (s *Server) IncrementInactiveCount(req *models.IncrementInactiveCountReq) (*models.ChangeInactiveCountResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	// find the activity in database
	var activity gormModel.Activity
	var user gormModel.User

	// if activity cannot be found by given ID, return error
	if result := s.Database.First(&activity, req.ActivityId); result.Error != nil {
		return nil, ErrActivityNotFound
	}
	if result := s.Database.First(&user, req.UserId); result.Error != nil {
		return nil, ErrUserNotFound
	}

	reportHistory := &gormModel.ReportHistory{
		UserId:     req.UserId,
		ActivityId: req.ActivityId,
		Reason:     req.Reason,
	}
	if result := s.Database.Create(&reportHistory); result.Error != nil {
		return nil, ErrAlreadyReported
	}

	activity.InactiveCount++
	if activity.InactiveCount >= InvalidThreshold {
		activity.InactiveFlag = true
	}

	if result := s.Database.Save(&activity); result.Error != nil {
		fmt.Println("inactivate_activity err: ", result.Error) // TODO: write to log instead
		return nil, result.Error
	}

	return &models.ChangeInactiveCountResp{
		ActivityId:    activity.ID,
		InactiveCount: activity.InactiveCount,
		InactiveFlag:  activity.InactiveFlag,
		UpdatedAt:     activity.UpdatedAt,
	}, nil
}

func (s *Server) DecrementInactiveCount(req *models.DecrementInactiveCountReq) (*models.ChangeInactiveCountResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	// find the activity in database
	var activity gormModel.Activity
	var user gormModel.User

	// if activity cannot be found by given ID, return error
	if result := s.Database.First(&activity, req.ActivityId); result.Error != nil {
		return nil, ErrActivityNotFound
	}
	if result := s.Database.First(&user, req.UserId); result.Error != nil {
		return nil, ErrUserNotFound
	}
	// delete history
	var reportHistory gormModel.ReportHistory
	if result := s.Database.Where("user_id=? AND activity_id = ?", req.UserId, req.ActivityId).Find(&reportHistory); result.Error != nil || result.RowsAffected == 0 {
		return nil, ErrReportNotFound
	}
	if err := s.Database.Unscoped().Model(&activity).Association("UserReports").Delete(&reportHistory); err != nil {
		fmt.Println(err.Error())
		return nil, ErrDatabase
	}
	// if succeeded, delete activity
	activity.InactiveCount--
	if activity.InactiveCount < InvalidThreshold {
		activity.InactiveFlag = false
	}

	if result := s.Database.Save(&activity); result.Error != nil {
		fmt.Println("inactivate_activity err: ", result.Error) // TODO: write to log instead
		return nil, ErrDatabase
	}

	return &models.ChangeInactiveCountResp{
		ActivityId:    activity.ID,
		InactiveCount: activity.InactiveCount,
		InactiveFlag:  activity.InactiveFlag,
		UpdatedAt:     activity.UpdatedAt,
	}, nil
}

func (s *Server) CheckInactiveFlag(req *models.HasUserInactivatedReq) (*models.HasUserInactivatedResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	// find the activity in database
	var reportHistory gormModel.ReportHistory

	if result := s.Database.Where("user_id=? AND activity_id = ?", req.UserId, req.ActivityId).Find(&reportHistory); result.Error != nil || result.RowsAffected == 0 {
		return &models.HasUserInactivatedResp{
			UpdatedAt: time.Now(),
			Reported:  false,
		}, nil
	}

	return &models.HasUserInactivatedResp{
		UpdatedAt: time.Now(),
		Reported:  true,
	}, nil
}

func (s *Server) DeleteActivityImage(req *models.DeleteActivityImageReq) (*models.DeleteActivityImageResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	// find the activity in database
	var activity gormModel.Activity

	// if activity cannot be found by given ID, return error
	if result := s.Database.First(&activity, req.ActivityId); result.Error != nil || result.RowsAffected == 0 {
		return nil, ErrActivityNotFound
	}

	// see if the user id matches the activity's user id
	if activity.UserID != req.UserId {
		return nil, ErrInvalidUpdateUser
	}

	idx := SearchName(activity.ImageNames, req.ImageName)
	if idx >= len(activity.ImageNames) {
		return nil, ErrImageNoMatch
	}
	err := os.Remove(filepath.Join(ImageRoot, ActivityImageFolder, req.ImageName))
	if err != nil {
		fmt.Println("image delete unsuccessful, ", err)
		activity.ImageNames = RemoveName(activity.ImageNames, idx)
		if result := s.Database.Save(&activity); result.Error != nil {
			fmt.Println("delete_image err: ", result.Error) // TODO: write to log instead
			return nil, ErrDatabase
		}
		return nil, ErrImageNotFound
	}

	activity.ImageNames = RemoveName(activity.ImageNames, idx)

	if result := s.Database.Save(&activity); result.Error != nil {
		fmt.Println("delete_image err: ", result.Error) // TODO: write to log instead
		return nil, result.Error
	}
	return &models.DeleteActivityImageResp{
		ActivityId: activity.ID,
		DeletedAt:  activity.UpdatedAt,
	}, nil

}

func (s *Server) UpdateReview(req *models.UpdateReviewReq) (*models.GetActivityResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	var review gormModel.Review
	// find review in the database by review id
	if result := s.Database.Where("id=? AND user_id=? AND activity_id=?", req.ReviewId, req.UserId, req.ActivityId).Find(&review); result.RowsAffected == 0 {
		return nil, ErrReviewNotFound
	}

	var activity gormModel.Activity
	// if activity cannot be found by given ID, return error
	if result := s.Database.Where("id=?", req.ActivityId).Preload("Reviews").Find(&activity); result.Error != nil || result.RowsAffected == 0 {
		return nil, ErrActivityNotFound
	}

	// user update
	var newAvg float32
	if activity.ReviewCounts-1 == 0 {
		newAvg = 0
	} else {
		newAvg = (float32(activity.ReviewCounts)*activity.AverageRating - review.Rating) / float32(activity.ReviewCounts-1)
	}
	activity.ReviewCounts--
	activity.AverageRating = newAvg
	if req.Delete == true {
		err := s.Database.Model(&activity).Association("Reviews").Delete(&review)
		if err != nil {
			return nil, ErrDatabase
		}
		s.Database.Unscoped().Delete(&review)
	} else {
		newAvg = (float32(activity.ReviewCounts)*activity.AverageRating + req.NewRating) / float32(activity.ReviewCounts+1)
		activity.ReviewCounts++
		activity.AverageRating = newAvg
		review.Title = req.Title
		review.Description = req.Description
		review.Rating = req.NewRating
		if result := s.Database.Save(&review); result.Error != nil {
			fmt.Println("create_review err: ", result.Error) // TODO: write to log instead
			return nil, ErrDatabase
		}
	}

	if result := s.Database.Save(&activity); result.Error != nil {
		fmt.Println("save activity err: ", result.Error) // TODO: write to log instead
		return nil, ErrDatabase
	}

	// if activity cannot be found by given ID, return error
	if result := s.Database.Where("id=?", req.ActivityId).Preload("Reviews").Find(&activity); result.Error != nil || result.RowsAffected == 0 {
		return nil, ErrActivityNotFound
	}

	return ParseActivity(activity), nil
}

// ValidateFile onwards are utility functions

func ValidateFile(fileHeader *multipart.FileHeader) (bool, error) {
	// open the uploaded file

	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("Cannot open file", err)
		return false, err
	}

	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Cannot close file", err)
			panic(err) //TODO: panic or send message?
		}
	}()

	buf := make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		fmt.Println("Cannot read file to buff", err)
		return false, err
	}

	filetype := http.DetectContentType(buf)
	switch filetype {
	case "image/jpeg", "image/jpg", "image/gif", "image/png": //"application/pdf" //TODO: allow PDF?
		fmt.Println("received image of type: " + filetype)
		return true, nil
	default:
		return false, ErrUnknownFileType
	}
}

func SaveFile(image *multipart.FileHeader, c *gin.Context, subDirectory string) (string, string, error) {
	_, err := ValidateFile(image)
	if err != nil {
		return "", "", err
	}
	if image == nil {
		return "", "", errors.New("Cannot find file")
	}
	// create subfolder if it doesn't exist
	fileDirectory := filepath.Join(ImageRoot, subDirectory)
	if _, err := os.Stat(fileDirectory); errors.Is(err, os.ErrNotExist) {
		mkdirErr := os.MkdirAll(fileDirectory, os.ModePerm) // define different file access
		if mkdirErr != nil {
			fmt.Println(mkdirErr) // TODO: log
		} else {
			fmt.Printf("Created %s at %s\n", fileDirectory, ImageRoot)
		}
	}
	uniqueImgName := uuid.NewString() + filepath.Ext(image.Filename)
	fPath := filepath.Join(fileDirectory, uniqueImgName)
	if _, err := os.Create(fPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "", "", err
	}
	if err := c.SaveUploadedFile(image, fPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "", "", err
	}
	return uniqueImgName, fPath, nil
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

func SearchName(s []string, name string) int {
	i := 0
	for ; i < len(s); i++ {
		if s[i] == name {
			break
		}
	}
	return i
}

// RemoveName of image by replacing the to be removed value with the last element. assuming image order doesn't matter
func RemoveName(s []string, i int) []string {
	if i > len(s) {
		return s
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func ParseReviewList(reviewList []gormModel.Review) []*models.Review {
	parsedReview := make([]*models.Review, 0)
	for _, review := range reviewList {
		parsedReview = append(parsedReview, &models.Review{
			ID:          review.ID,
			UserId:      review.UserId,
			ActivityId:  review.ActivityId,
			Title:       review.Title,
			Description: review.Description,
			Rating:      review.Rating,
		})
	}
	return parsedReview
}

func PackCreateOpeningTimes(createForm *models.CreateActivityForm) []int32 {
	var opening []int32
	opening = append(opening, int32(getValidTime(createForm.SunOpeningTime)),
		int32(getValidTime(createForm.MonOpeningTime)),
		int32(getValidTime(createForm.TueOpeningTime)), int32(getValidTime(createForm.WedOpeningTime)),
		int32(getValidTime(createForm.ThurOpeningTime)), int32(getValidTime(createForm.FriOpeningTime)),
		int32(getValidTime(createForm.SatOpeningTime)), int32(getValidTime(createForm.SunClosingTime)),
		int32(getValidTime(createForm.MonClosingTime)), int32(getValidTime(createForm.TueClosingTime)),
		int32(getValidTime(createForm.WedClosingTime)), int32(getValidTime(createForm.ThurClosingTime)),
		int32(getValidTime(createForm.FriClosingTime)), int32(getValidTime(createForm.SatClosingTime)),
	)
	return opening
}

func PackUpdateOpeningTimes(updateReq *models.UpdateActivityForm) []int32 {
	var opening []int32
	opening = append(opening, int32(getValidTime(updateReq.SunOpeningTime)),
		int32(getValidTime(updateReq.MonOpeningTime)),
		int32(getValidTime(updateReq.TueOpeningTime)), int32(getValidTime(updateReq.WedOpeningTime)),
		int32(getValidTime(updateReq.ThurOpeningTime)), int32(getValidTime(updateReq.FriOpeningTime)),
		int32(getValidTime(updateReq.SatOpeningTime)), int32(getValidTime(updateReq.SunClosingTime)),
		int32(getValidTime(updateReq.MonClosingTime)), int32(getValidTime(updateReq.TueClosingTime)),
		int32(getValidTime(updateReq.WedClosingTime)), int32(getValidTime(updateReq.ThurClosingTime)),
		int32(getValidTime(updateReq.FriClosingTime)), int32(getValidTime(updateReq.SatClosingTime)))
	return opening
}

func ParseActivity(activity gormModel.Activity) *models.GetActivityResp {
	return &models.GetActivityResp{
		ActivityId:  activity.ID,
		Title:       activity.Title,
		Rating:      float32(Round(float64(activity.AverageRating), 0.05)),
		Paid:        activity.Paid,
		Categories:  activity.Categories,
		Description: activity.Description,
		Longitude:   activity.Longitude,
		Latitude:    activity.Latitude,
		ImageNames:  activity.ImageNames,

		SunOpeningTime:  int(activity.OpeningTimes[0]),
		MonOpeningTime:  int(activity.OpeningTimes[1]),
		TueOpeningTime:  int(activity.OpeningTimes[2]),
		WedOpeningTime:  int(activity.OpeningTimes[3]),
		ThurOpeningTime: int(activity.OpeningTimes[4]),
		FriOpeningTime:  int(activity.OpeningTimes[5]),
		SatOpeningTime:  int(activity.OpeningTimes[6]),

		SunClosingTime:  int(activity.OpeningTimes[7]),
		MonClosingTime:  int(activity.OpeningTimes[8]),
		TueClosingTime:  int(activity.OpeningTimes[9]),
		WedClosingTime:  int(activity.OpeningTimes[10]),
		ThurClosingTime: int(activity.OpeningTimes[11]),
		FriClosingTime:  int(activity.OpeningTimes[12]),
		SatClosingTime:  int(activity.OpeningTimes[13]),

		InactiveCount: activity.InactiveCount,
		InactiveFlag:  activity.InactiveFlag,
		ReviewCounts:  activity.ReviewCounts,
		ReviewsList:   ParseReviewList(activity.Reviews),
		CreatedAt:     activity.CreatedAt,
	}
}
