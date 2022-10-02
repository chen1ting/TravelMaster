package server

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"

	"github.com/google/uuid"

	gormModel "github.com/chen1ting/TravelMaster/internal/models/gorm"

	"github.com/chen1ting/TravelMaster/internal/models"
)

var (
	ErrBadRequest            = errors.New("bad request")
	ErrMissingUserInfo       = errors.New("eight email, username, or hashed password missing")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrInvalidLogin          = errors.New("invalid login")
	ErrActivityAlreadyExists = errors.New("an activity with the same title already exists")
	ErrNullTitle             = errors.New("title cannot be empty")
	ErrInvalidActivityID     = errors.New("activity id doesn't exist")
	ErrInvalidCreateUser     = errors.New("user id doesn't exists")
	ErrInvalidUpdateUser     = errors.New("user id doesn't match the activity's user id")
	ErrNoSearchFail          = errors.New("searchName failed")
	ErrParsingResultFail     = errors.New("cannot parse result")
	ErrUnknownFileType       = errors.New("unknown file type uploaded")
	ErrImageNoMatch          = errors.New("image not found in the list of the activity")
	ErrImageNotFound         = errors.New("image not found on server, removed file name in the database")
	CWD, _                   = os.Getwd()
	ImageRoot                = filepath.Join(CWD, "assets")
	ActivityImageFolder      = "activity_images"
	AvatarFolder             = "avatars"
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
		Interests:  form.Interests,
		AvatarName: uniqueImgName,
	}

	if result := s.Database.Model(&user).Create(&user); result.Error != nil {
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

func (s *Server) GenerateItinerary(ctx context.Context, req *models.GenerateItineraryRequest) (*models.GenerateItineraryResponse, error) {
	// TODO: for now hardcoded resp from 29/9/2022, 3:00 PM to 30/9/2022, 11:00 PM
	return &models.GenerateItineraryResponse{
		GeneratedItinerary: &models.Itinerary{
			Id: 100,
			StartTime: 1664434800,
			EndTime: 1664550000,
			NumberOfSegments: 3,
			Segments: []*models.Segment{
				{
					StartTime: 1664434800, // 29/9/2022, 3:00 PM
					EndTime: 1664442000, // 29/9/2022, 5:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 1,
						Name: "Sample activity 1",
						Description: "This is a sample and test activity...",
						AverageRating: 3.5,
						Categories: []string{"Fun"},
						ImageUrl: "https://images.pexels.com/photos/457882/pexels-photo-457882.jpeg?auto=compress&cs=tinysrgb&w=800",
					},
				},
				{
					StartTime: 1664449200, // 30/9/2022, 7:00 PM
					EndTime: 1664456400, // 29/9/2022, 9:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 5,
						Name: "Restaurant ABC",
						Description: "This is an activity for your dinner",
						AverageRating: 4,
						Categories: []string{"Food"},
						ImageUrl: "https://media.cntraveler.com/photos/61eae2a9fe18edcbd885cb01/5:4/w_3790,h_3032,c_limit/Seychelles_GettyImages-1169388113.jpg",
					},
				},
				{
					StartTime: 1664456400, // 29/9/2022, 9:00 PM
					EndTime: 1664463600, // 29/9/2022, 11:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 2,
						Name: "Sample activity 2",
						Description: "This is a sample and test activity...",
						AverageRating: 4,
						Categories: []string{"Something Long"},
						ImageUrl: "https://media.istockphoto.com/photos/tropical-white-sand-beach-with-coco-palms-picture-id1181563943?k=20&m=1181563943&s=612x612&w=0&h=r46MQMvFnvrzzTfjVmvZED5nZyTmAYwISDvkdtM2i2A=",
					},
				},
				{
					StartTime: 1664496000, // 30/9/2022, 8:00 AM
					EndTime: 1664503200, // 30/9/2022, 10:00 AM
					ActivitySummary: models.ActivitySummary{
						Id: 200,
						Name: "Breakie Time",
						Description: "This is for breakfast.",
						AverageRating: 4,
						Categories: []string{"Food"},
						ImageUrl: "https://static.thehoneycombers.com/wp-content/uploads/sites/4/2022/04/Sundays-Beach-Club-in-Uluwatu-Bali-Indonesia.jpeg",
					},
				},
				{
					StartTime: 1664503200, // 30/9/2022, 10:00 AM
					EndTime: 1664510400, // 30/9/2022, 12:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 3,
						Name: "Sample activity 3",
						Description: "This is a sample and test activity...",
						AverageRating: 4,
						Categories: []string{"Something else", "Another thing", "Fun"},
						ImageUrl: "https://image.shutterstock.com/image-photo/chairs-umbrella-palm-beach-tropical-260nw-559599520.jpg",
					},
				},
				{
					StartTime: 1664510400, // 30/9/2022, 12:00 PM
					EndTime: 1664517600, // 30/9/2022, 2:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 10,
						Name: "Lunch time bby",
						Description: "This is for lunch",
						AverageRating: 4.5,
						Categories: []string{"Food"},
						ImageUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRxkErgnfKaYHK1adHcY02d7f_B7sD0mwSLMg&usqp=CAU",
					},
				},
				{
					StartTime: 1664517600, // 30/9/2022, 2:00 PM
					EndTime: 1664524800, // 30/9/2022, 4:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 4,
						Name: "Sample activity 4",
						Description: "This is a sample and test activity...",
						AverageRating: 4,
						Categories: []string{"Something Long", "ABC123", "another thing"},
						ImageUrl: "https://media.istockphoto.com/photos/tropical-beach-palm-trees-sea-wave-and-white-sand-picture-id1300296030?b=1&k=20&m=1300296030&s=170667a&w=0&h=w1s7kmN2TH7O326d263Cs-E44teA1hy6u29UIVf_z1w=",
					},
				},
				{
					StartTime: 1664532000, // 30/9/2022, 6:00 PM
					EndTime: 1664535600, // 30/9/2022, 7:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 5,
						Name: "Sample activity 5",
						Description: "This is a sample and test activity...",
						AverageRating: 2,
						Categories: []string{"Something Long", "ABC123", "another thing"},
						ImageUrl: "https://i.insider.com/5bfec49248eb12058423acf7?width=700",
					},
				},
				{
					StartTime: 1664535600, // 30/9/2022, 7:00 PM
					EndTime: 1664542800, // 30/9/2022, 9:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 5,
						Name: "Restaurant ABC",
						Description: "This is an activity for your dinner",
						AverageRating: 4,
						Categories: []string{"Food"},
						ImageUrl: "https://media.cntraveler.com/photos/61eae2a9fe18edcbd885cb01/5:4/w_3790,h_3032,c_limit/Seychelles_GettyImages-1169388113.jpg",
					},
				},
				{
					StartTime: 1664542800, // 30/9/2022, 9:00 PM
					EndTime: 1664550000, // 30/9/2022, 11:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 6,
						Name: "Sample activity 6",
						Description: "This is a sample and test activity...",
						AverageRating: 3,
						Categories: []string{"Something Long", "ABC123", "another thing"},
						ImageUrl: "https://www.travelandleisure.com/thmb/mUQhPpdcuEfnsprkLEyuT4BWK_M=/1800x1012/smart/filters:no_upscale()/saud-beach-luzon-philippines-WRLDBEACH0421-15e2c368e7ad4495be803bd60cafa379.jpg",
					},
				},
			},
		},
	}, nil
}


func (s *Server) GetItinerary(ctx context.Context, req *models.GetItineraryRequest) (*models.GetItineraryResponse, error) {
	// TODO: for now hardcoded resp from 29/9/2022, 3:00 PM to 30/9/2022, 11:00 PM
	return &models.GetItineraryResponse{
		Itinerary: &models.Itinerary{
			Id: 100,
			StartTime: 1664434800,
			EndTime: 1664550000,
			NumberOfSegments: 3,
			Segments: []*models.Segment{
				{
					StartTime: 1664434800, // 29/9/2022, 3:00 PM
					EndTime: 1664442000, // 29/9/2022, 5:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 1,
						Name: "Sample activity 1",
						Description: "This is a sample and test activity...",
						AverageRating: 3.5,
						Categories: []string{"Fun"},
						ImageUrl: "https://images.pexels.com/photos/457882/pexels-photo-457882.jpeg?auto=compress&cs=tinysrgb&w=800",
					},
				},
				{
					StartTime: 1664449200, // 30/9/2022, 7:00 PM
					EndTime: 1664456400, // 29/9/2022, 9:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 5,
						Name: "Restaurant ABC",
						Description: "This is an activity for your dinner",
						AverageRating: 4,
						Categories: []string{"Food"},
						ImageUrl: "https://media.cntraveler.com/photos/61eae2a9fe18edcbd885cb01/5:4/w_3790,h_3032,c_limit/Seychelles_GettyImages-1169388113.jpg",
					},
				},
				{
					StartTime: 1664456400, // 29/9/2022, 9:00 PM
					EndTime: 1664463600, // 29/9/2022, 11:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 2,
						Name: "Sample activity 2",
						Description: "This is a sample and test activity...",
						AverageRating: 4,
						Categories: []string{"Something Long"},
						ImageUrl: "https://media.istockphoto.com/photos/tropical-white-sand-beach-with-coco-palms-picture-id1181563943?k=20&m=1181563943&s=612x612&w=0&h=r46MQMvFnvrzzTfjVmvZED5nZyTmAYwISDvkdtM2i2A=",
					},
				},
				{
					StartTime: 1664496000, // 30/9/2022, 8:00 AM
					EndTime: 1664503200, // 30/9/2022, 10:00 AM
					ActivitySummary: models.ActivitySummary{
						Id: 200,
						Name: "Breakie Time",
						Description: "This is for breakfast.",
						AverageRating: 4,
						Categories: []string{"Food"},
						ImageUrl: "https://static.thehoneycombers.com/wp-content/uploads/sites/4/2022/04/Sundays-Beach-Club-in-Uluwatu-Bali-Indonesia.jpeg",
					},
				},
				{
					StartTime: 1664503200, // 30/9/2022, 10:00 AM
					EndTime: 1664510400, // 30/9/2022, 12:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 3,
						Name: "Sample activity 3",
						Description: "This is a sample and test activity...",
						AverageRating: 4,
						Categories: []string{"Something else", "Another thing", "Fun"},
						ImageUrl: "https://image.shutterstock.com/image-photo/chairs-umbrella-palm-beach-tropical-260nw-559599520.jpg",
					},
				},
				{
					StartTime: 1664510400, // 30/9/2022, 12:00 PM
					EndTime: 1664517600, // 30/9/2022, 2:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 10,
						Name: "Lunch time bby",
						Description: "This is for lunch",
						AverageRating: 4.5,
						Categories: []string{"Food"},
						ImageUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRxkErgnfKaYHK1adHcY02d7f_B7sD0mwSLMg&usqp=CAU",
					},
				},
				{
					StartTime: 1664517600, // 30/9/2022, 2:00 PM
					EndTime: 1664524800, // 30/9/2022, 4:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 4,
						Name: "Sample activity 4",
						Description: "This is a sample and test activity...",
						AverageRating: 4,
						Categories: []string{"Something Long", "ABC123", "another thing"},
						ImageUrl: "https://media.istockphoto.com/photos/tropical-beach-palm-trees-sea-wave-and-white-sand-picture-id1300296030?b=1&k=20&m=1300296030&s=170667a&w=0&h=w1s7kmN2TH7O326d263Cs-E44teA1hy6u29UIVf_z1w=",
					},
				},
				{
					StartTime: 1664532000, // 30/9/2022, 6:00 PM
					EndTime: 1664535600, // 30/9/2022, 7:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 5,
						Name: "Sample activity 5",
						Description: "This is a sample and test activity...",
						AverageRating: 2,
						Categories: []string{"Something Long", "ABC123", "another thing"},
						ImageUrl: "https://i.insider.com/5bfec49248eb12058423acf7?width=700",
					},
				},
				{
					StartTime: 1664535600, // 30/9/2022, 7:00 PM
					EndTime: 1664542800, // 30/9/2022, 9:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 5,
						Name: "Restaurant ABC",
						Description: "This is an activity for your dinner",
						AverageRating: 4,
						Categories: []string{"Food"},
						ImageUrl: "https://media.cntraveler.com/photos/61eae2a9fe18edcbd885cb01/5:4/w_3790,h_3032,c_limit/Seychelles_GettyImages-1169388113.jpg",
					},
				},
				{
					StartTime: 1664542800, // 30/9/2022, 9:00 PM
					EndTime: 1664550000, // 30/9/2022, 11:00 PM
					ActivitySummary: models.ActivitySummary{
						Id: 6,
						Name: "Sample activity 6",
						Description: "This is a sample and test activity...",
						AverageRating: 3,
						Categories: []string{"Something Long", "ABC123", "another thing"},
						ImageUrl: "https://www.travelandleisure.com/thmb/mUQhPpdcuEfnsprkLEyuT4BWK_M=/1800x1012/smart/filters:no_upscale()/saud-beach-luzon-philippines-WRLDBEACH0421-15e2c368e7ad4495be803bd60cafa379.jpg",
					},
				},
			},
		},
	}, nil
}

func (s *Server) GetActivitiesByFilter(ctx context.Context, req * models.GetActivitiesByFilterRequest) (*models.GetActivitiesByFilterResponse, error) {
	return &models.GetActivitiesByFilterResponse{
		NumOfResults: 3,
		Activities: []*models.ActivitySummary{
			{
				Id: 1,
				Name: "Sample test 1",
				Description: "Sample description of this activity...",
				AverageRating: 3.5,
				Categories: []string{"Fun"},
				ImageUrl: "https://visitbeaches.org/static/media/beach-stock.a6ea40bc.jpeg",
			},
			{
				Id: 2,
				Name: "Sample test 2",
				Description: "Sample description of this activity...",
				AverageRating: 3,
				Categories: []string{"Fun", "Something else"},
				ImageUrl: "https://www.ucdavis.edu/sites/default/files/styles/ucd_panoramic_image/public/media/images/beaches-near-uc-davis.jpg?h=8e58fdb5&itok=0D79HHcC",
			},
			{
				Id: 3,
				Name: "Sample test 3",
				Description: "Sample description of this activity...",
				AverageRating: 5,
				Categories: []string{"Fun", "Something else"},
				ImageUrl: "https://visitbeaches.org/static/media/beach-stock.a6ea40bc.jpeg",
			},
		},
	}, nil
}

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
	opening = append(opening, int32(getValidTime(createForm.MonOpeningTime)),
		int32(getValidTime(createForm.TueOpeningTime)), int32(getValidTime(createForm.WedOpeningTime)),
		int32(getValidTime(createForm.ThurOpeningTime)), int32(getValidTime(createForm.FriOpeningTime)),
		int32(getValidTime(createForm.SatOpeningTime)), int32(getValidTime(createForm.SunOpeningTime)),
		int32(getValidTime(createForm.MonClosingTime)), int32(getValidTime(createForm.TueClosingTime)),
		int32(getValidTime(createForm.WedClosingTime)), int32(getValidTime(createForm.ThurClosingTime)),
		int32(getValidTime(createForm.FriClosingTime)), int32(getValidTime(createForm.SatClosingTime)),
		int32(getValidTime(createForm.SunClosingTime)))
	return opening
}

func packUpdateOpeningTimes(updateReq *models.UpdateActivityForm) []int32 {
	var opening []int32
	opening = append(opening, int32(getValidTime(updateReq.MonOpeningTime)),
		int32(getValidTime(updateReq.TueOpeningTime)), int32(getValidTime(updateReq.WedOpeningTime)),
		int32(getValidTime(updateReq.ThurOpeningTime)), int32(getValidTime(updateReq.FriOpeningTime)),
		int32(getValidTime(updateReq.SatOpeningTime)), int32(getValidTime(updateReq.SunOpeningTime)),
		int32(getValidTime(updateReq.MonClosingTime)), int32(getValidTime(updateReq.TueClosingTime)),
		int32(getValidTime(updateReq.WedClosingTime)), int32(getValidTime(updateReq.ThurClosingTime)),
		int32(getValidTime(updateReq.FriClosingTime)), int32(getValidTime(updateReq.SatClosingTime)),
		int32(getValidTime(updateReq.SunClosingTime)))
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

func (s *Server) CreateActivity(form *models.CreateActivityForm, c *gin.Context) (*models.CreateActivityResp, error) {
	if form == nil {
		return nil, ErrBadRequest
	}

	var user gormModel.User
	if result := s.Database.First(&user, form.UserId); result.Error != nil {
		return nil, ErrInvalidCreateUser
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
		UserID:        form.UserId,
		Title:         form.Title,
		AverageRating: form.Rating,
		Paid:          form.Paid,
		Category:      form.Category,
		Description:   form.Description,
		Longitude:     form.Longitude,
		Latitude:      form.Latitude,
		OpeningTimes:  packCreateOpeningTimes(form),
		ImageNames:    imgNames,

		// system settings
		InactiveCount: 0,
		InactiveFlag:  false,
		ReviewCounts:  0,
		ReviewIds:     "",
	}

	if result := s.Database.Model(&activity).Create(&activity); result.Error != nil {
		fmt.Println("create_activity err: ", result.Error) // TODO: write to log instead
		// if result cannot be saved, removeName all saved images
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
		return nil, ErrInvalidActivityID
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
	activity.AverageRating = form.Rating
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
	if result := s.Database.First(&activity, req.ActivityId); result.Error != nil {
		return nil, ErrInvalidActivityID
	}

	return &models.GetActivityResp{
		ActivityId:  activity.ID,
		Title:       activity.Title,
		Rating:      activity.AverageRating,
		Paid:        activity.Paid,
		Category:    activity.Category,
		Description: activity.Description,
		Longitude:   activity.Longitude,
		Latitude:    activity.Latitude,
		ImageNames:  activity.ImageNames,

		MonOpeningTime:  int(activity.OpeningTimes[0]),
		TueOpeningTime:  int(activity.OpeningTimes[1]),
		WedOpeningTime:  int(activity.OpeningTimes[2]),
		ThurOpeningTime: int(activity.OpeningTimes[3]),
		FriOpeningTime:  int(activity.OpeningTimes[4]),
		SatOpeningTime:  int(activity.OpeningTimes[5]),
		SunOpeningTime:  int(activity.OpeningTimes[6]),
		MonClosingTime:  int(activity.OpeningTimes[7]),
		TueClosingTime:  int(activity.OpeningTimes[8]),
		WedClosingTime:  int(activity.OpeningTimes[9]),
		ThurClosingTime: int(activity.OpeningTimes[10]),
		FriClosingTime:  int(activity.OpeningTimes[11]),
		SatClosingTime:  int(activity.OpeningTimes[12]),
		SunClosingTime:  int(activity.OpeningTimes[13]),

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

func searchName(s []string, name string) int {
	i := 0
	for ; i < len(s); i++ {
		if s[i] == name {
			break
		}
	}
	return i
}

// assuming image order doesn't matter
func removeName(s []string, i int) []string {
	if i > len(s) {
		return s
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (s *Server) DeleteActivityImage(req *models.DeleteActivityImageReq) (*models.DeleteActivityImageResp, error) {
	if req == nil {
		return nil, ErrBadRequest
	}

	// find the activity in database
	var activity gormModel.Activity

	// if activity cannot be found by given ID, return error
	if result := s.Database.First(&activity, req.ActivityId); result.Error != nil || result.RowsAffected == 0 {
		return nil, ErrInvalidActivityID
	}

	// see if the user id matches the activity's user id
	if activity.UserID != req.UserId {
		return nil, ErrInvalidUpdateUser
	}

	idx := searchName(activity.ImageNames, req.ImageName)
	fmt.Println(idx)
	if idx >= len(activity.ImageNames) {
		return nil, ErrImageNoMatch
	}
	err := os.Remove(filepath.Join(ImageRoot, ActivityImageFolder, req.ImageName))
	if err != nil {
		fmt.Println("image delete unsuccessful, ", err)
		activity.ImageNames = removeName(activity.ImageNames, idx)
		if result := s.Database.Save(&activity); result.Error != nil {
			fmt.Println("delete_image err: ", result.Error) // TODO: write to log instead
			return nil, result.Error
		}
		return nil, ErrImageNotFound
	}

	activity.ImageNames = removeName(activity.ImageNames, idx)

	if result := s.Database.Save(&activity); result.Error != nil {
		fmt.Println("delete_image err: ", result.Error) // TODO: write to log instead
		return nil, result.Error
	}
	return &models.DeleteActivityImageResp{
		ActivityId: activity.ID,
		DeletedAt:  activity.UpdatedAt,
	}, nil

}
