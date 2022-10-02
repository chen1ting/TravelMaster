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

	// attempt to save user to DB
	user := gormModel.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.HashedPassword,
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