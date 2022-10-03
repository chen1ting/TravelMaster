package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/google/uuid"

	gormModel "github.com/chen1ting/TravelMaster/internal/models/gorm"

	"github.com/chen1ting/TravelMaster/internal/models"
)

var (
	ErrBadRequest            = errors.New("bad request")
	ErrMissingUserInfo       = errors.New("missing necessary user information")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrInvalidLogin          = errors.New("invalid login")
	ErrActivityAlreadyExists = errors.New("an activity with the same title already exists")
	ErrUserNotFound          = errors.New("user id doesn't exists")
	ErrNullTitle             = errors.New("title cannot be empty")
	ErrNullReview            = errors.New("review content cannot be empty")
	ErrActivityNotFound      = errors.New("activity id doesn't exist")
	ErrReviewNotFound        = errors.New("review id doesn't exist")
	ErrInvalidUpdateUser     = errors.New("user id doesn't match the activity's user id")
	ErrNoSearchFail          = errors.New("search name failed")
	ErrUnknownFileType       = errors.New("unknown file type uploaded")
	ErrImageNoMatch          = errors.New("image not found in the list of the activity")
	ErrImageNotFound         = errors.New("image not found on server, removed file name in the database")
	// database error code reference https://github.com/jackc/pgerrcode/blob/master/errcode.go
	CWD, _              = os.Getwd()
	ImageRoot           = filepath.Join(CWD, "assets")
	ActivityImageFolder = "activity_images"
	AvatarFolder        = "avatars"
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

	//
	uniqueImgName, fpath, saveErr := SaveFile(form.Avatar, c, AvatarFolder)
	if saveErr != nil {
		fmt.Println(saveErr) //TODO: log instead
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
			fmt.Println("sign_up have error deleting avatar: ", err)
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
	if result := s.Database.Where("id=?", req.UserId).Preload("Activities").Preload("Reviews").Find(&user); result.Error != nil {
		//if result := s.Database.First(&user, req.UserId); result.Error != nil {
		return nil, ErrUserNotFound
	}
	return &models.GetProfileResp{
		User:        user,
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
	avatarName, avatarPath, saveErr := SaveFile(form.Avatar, c, ActivityImageFolder)
	if saveErr != nil {
		return nil, saveErr
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
		Category:     form.Category,
		Description:  form.Description,
		Longitude:    form.Longitude,
		Latitude:     form.Latitude,
		OpeningTimes: packCreateOpeningTimes(form),
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
	activity.Category = form.Category
	activity.Description = form.Description
	activity.Longitude = form.Longitude
	activity.Latitude = form.Latitude
	activity.ImageNames = imgNames
	activity.OpeningTimes = packUpdateOpeningTimes(form)

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

func (s *Server) GetActivity(req *models.GetActivityReq) (*models.GetActivityResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}
	var activity gormModel.Activity

	// if activity cannot be found by given ID, return error
	if result := s.Database.Where("id=?", req.ActivityId).Preload("Reviews").Find(&activity); result.Error != nil {
		return nil, ErrActivityNotFound
	}

	return &models.GetActivityResp{
		Activity:    activity,
		RetrievedAt: time.Now(),
	}, nil
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
	return &models.SearchActivityResp{
		Activities:   activities,
		ResultNumber: result.RowsAffected,
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
		return nil, ErrActivityNotFound
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
	fmt.Println(idx)
	if idx >= len(activity.ImageNames) {
		return nil, ErrImageNoMatch
	}
	err := os.Remove(filepath.Join(ImageRoot, ActivityImageFolder, req.ImageName))
	if err != nil {
		fmt.Println("image delete unsuccessful, ", err)
		activity.ImageNames = RemoveName(activity.ImageNames, idx)
		if result := s.Database.Save(&activity); result.Error != nil {
			fmt.Println("delete_image err: ", result.Error) // TODO: write to log instead
			return nil, result.Error
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

func (s *Server) CreateReview(req *models.CreateReviewReq) (*models.CreateReviewResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	if req.Review == "" {
		return nil, ErrNullReview
	}

	var user gormModel.User
	// find the user in database
	if result := s.Database.First(&user, req.UserId); result.Error != nil {
		return nil, ErrUserNotFound
	}

	var activity gormModel.Activity
	// if activity cannot be found by given ID, return error
	if result := s.Database.First(&activity, req.ActivityId); result.Error != nil || result.RowsAffected == 0 {
		return nil, ErrActivityNotFound
	}

	review := gormModel.Review{
		UserId:     req.UserId,
		ActivityId: req.ActivityId,
		Review:     req.Review,
		Rating:     req.Rating,
	}
	if result := s.Database.Save(&review); result.Error != nil {
		fmt.Println("create_review err: ", result.Error) // TODO: write to log instead
		return nil, result.Error
	}

	if err := s.UpdateAverageReview(&activity, true, req.Rating); err != nil {
		return nil, err
	}

	return &models.CreateReviewResp{
		ReviewId:      review.ID,
		CreatedAt:     review.CreatedAt,
		ReviewCounts:  activity.ReviewCounts,
		AverageRating: activity.AvgReviewRating,
	}, nil
}

func (s *Server) UpdateReview(req *models.UpdateReviewReq) (*models.UpdateReviewResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	if req.Review == "" {
		return nil, ErrNullReview
	}

	var review gormModel.Review
	// find review in the database by review id
	if result := s.Database.Where("id=?", req.ReviewId, "user_id=?", req.UserId, "activity_id=?", req.ActivityId).Find(&review); result.Error != nil {
		return nil, ErrReviewNotFound
	}

	var activity gormModel.Activity
	// if activity cannot be found by given ID, return error
	if result := s.Database.First(&activity, req.ActivityId); result.Error != nil || result.RowsAffected == 0 {
		return nil, ErrActivityNotFound
	}

	if err := s.UpdateAverageReview(&activity, false, review.Rating); err != nil {
		return nil, err
	}

	if req.Delete {
		s.Database.Delete(&review)
		return &models.UpdateReviewResp{
			ReviewId:      req.ReviewId,
			UpdatedAt:     time.Now(),
			ReviewCounts:  activity.ReviewCounts,
			AverageRating: activity.AvgReviewRating,
		}, nil
	}

	if err := s.UpdateAverageReview(&activity, true, req.Rating); err != nil {
		return nil, err
	}

	review.Review = req.Review
	review.Rating = req.Rating
	if result := s.Database.Save(&review); result.Error != nil {
		fmt.Println("create_review err: ", result.Error) // TODO: write to log instead
		return nil, result.Error
	}

	return &models.UpdateReviewResp{
		ReviewId:      review.ID,
		UpdatedAt:     review.UpdatedAt,
		ReviewCounts:  activity.ReviewCounts,
		AverageRating: activity.AvgReviewRating,
	}, nil
}

// ValidateFile onwards are utility functions

func getValidTime(hhmm int) int {
	hour := hhmm / 100
	min := hhmm % 100
	if 0 <= hour && hour < 24 && 0 <= min && min < 60 {
		return hhmm
	}
	return -1
}

func packCreateOpeningTimes(createForm *models.CreateActivityForm) []int32 {
	var opening []int32
	opening = append(opening,
		int32(getValidTime(createForm.SunOpeningTime)), int32(getValidTime(createForm.SunClosingTime)),
		int32(getValidTime(createForm.MonOpeningTime)), int32(getValidTime(createForm.MonClosingTime)),
		int32(getValidTime(createForm.TueOpeningTime)), int32(getValidTime(createForm.TueClosingTime)),
		int32(getValidTime(createForm.WedOpeningTime)), int32(getValidTime(createForm.WedClosingTime)),
		int32(getValidTime(createForm.ThurOpeningTime)), int32(getValidTime(createForm.ThurClosingTime)),
		int32(getValidTime(createForm.FriOpeningTime)), int32(getValidTime(createForm.FriClosingTime)),
		int32(getValidTime(createForm.SatOpeningTime)), int32(getValidTime(createForm.SatClosingTime)))
	return opening
}

// to use getDay function, use idx/2
func packUpdateOpeningTimes(updateForm *models.UpdateActivityForm) []int32 {
	var opening []int32
	opening = append(opening,
		int32(getValidTime(updateForm.SunOpeningTime)), int32(getValidTime(updateForm.SunClosingTime)),
		int32(getValidTime(updateForm.MonOpeningTime)), int32(getValidTime(updateForm.MonClosingTime)),
		int32(getValidTime(updateForm.TueOpeningTime)), int32(getValidTime(updateForm.TueClosingTime)),
		int32(getValidTime(updateForm.WedOpeningTime)), int32(getValidTime(updateForm.WedClosingTime)),
		int32(getValidTime(updateForm.ThurOpeningTime)), int32(getValidTime(updateForm.ThurClosingTime)),
		int32(getValidTime(updateForm.FriOpeningTime)), int32(getValidTime(updateForm.FriClosingTime)),
		int32(getValidTime(updateForm.SatOpeningTime)), int32(getValidTime(updateForm.SatClosingTime)))
	return opening
}

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
		fmt.Println("unknown file type uploaded")
		return false, ErrUnknownFileType
	}
}

func SaveFile(image *multipart.FileHeader, c *gin.Context, subDirectory string) (string, string, error) {
	_, err := ValidateFile(image)
	if err != nil {
		return "", "", err
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

// assuming image order doesn't matter
func RemoveName(s []string, i int) []string {
	if i > len(s) {
		return s
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (s *Server) UpdateAverageReview(activity *gormModel.Activity, add bool, rating float32) error {
	//update the average rating of the activity if the new review is correctly saved
	var totalRating float32
	if add {
		totalRating = activity.AvgReviewRating * float32(activity.ReviewCounts)
		totalRating += rating
		activity.ReviewCounts++
	} else {
		totalRating = activity.AvgReviewRating * float32(activity.ReviewCounts)
		totalRating -= rating
		activity.ReviewCounts--
	}

	activity.AvgReviewRating = totalRating / float32(activity.ReviewCounts)
	if result := s.Database.Save(&activity); result.Error != nil {
		fmt.Println("create_review: ", result.Error) // TODO: write to log instead
		return result.Error
	}
	return nil
}
