package models

import (
	gormModel "github.com/chen1ting/TravelMaster/internal/models/gorm"
	"mime/multipart"
	"time"
)

type LoginReq struct {
	//assumptions: use username to login
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
}

type LoginResp struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	SessionToken string `json:"session_token"`
}

type LogoutReq struct {
	SessionToken string `json:"session_token"`
}

type SignupForm struct {
	Username       string                `form:"username"`
	HashedPassword string                `form:"hashed_password"`
	Email          string                `form:"email"`
	Avatar         *multipart.FileHeader `form:"avatar"`
}

type SignupResp struct {
	UserId     int64  `json:"user_id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	AvatarName string `json:"avatar_file_name"`

	SessionToken string `json:"session_token"`
}

type ValidateTokenReq struct {
	SessionToken string `json:"session_token"`
}

type ValidateTokenResp struct {
	Valid  bool  `json:"valid"`
	UserId int64 `json:"user_id"`
}

// TODO: allow change of usernames?
type UpdateProfileReq struct {
	UserId    int64    `json:"user_id"`
	AboutMe   string   `json:"about_me"`
	Interests []string `json:"interests"`
}

type UpdateProfileResp struct {
	UserId    int64     `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetProfileReq struct {
	UserId int64 `json:"user_id"`
}

type GetProfileResp struct {
	User        gormModel.User `json:"user"`
	RetrievedAt time.Time      `json:"retrieved_at"`
}

type UpdateAvatarForm struct {
	UserId int64                 `form:"user_id"`
	Delete bool                  `form:"delete"`
	Avatar *multipart.FileHeader `form:"avatar"`
}

type UpdateAvatarResp struct {
	UserId            int64     `json:"user_id"`
	UpdatedAt         time.Time `json:"updated_at"`
	NewAvatarFileName string    `json:"new_avtar_file_name"`
}

type CreateActivityForm struct {
	// Assumption: user token is already validated
	UserId   int64    `form:"user_id"`
	Title    string   `form:"title"`
	Rating   float32  `form:"rating_score"`
	Paid     bool     `form:"paid"`
	Category []string `form:"category"` // issue: form binding for string not working as expected
	// please send in a json style list of string
	Description string  `form:"description"`
	Longitude   float32 `form:"longitude"`
	Latitude    float32 `form:"latitude"`
	// Assumption: one image file upload at once
	Image []*multipart.FileHeader `form:"image"`

	// fields for opening & closing hours
	MonOpeningTime  int `form:"mon_opening_time"`
	MonClosingTime  int `form:"mon_closing_time"`
	TueOpeningTime  int `form:"tue_opening_time"`
	TueClosingTime  int `form:"tue_closing_time"`
	WedOpeningTime  int `form:"wed_opening_time"`
	WedClosingTime  int `form:"wed_closing_time"`
	ThurOpeningTime int `form:"thur_opening_time"`
	ThurClosingTime int `form:"thur_closing_time"`
	FriOpeningTime  int `form:"fri_opening_time"`
	FriClosingTime  int `form:"fri_closing_time"`
	SatOpeningTime  int `form:"sat_opening_time"`
	SatClosingTime  int `form:"sat_closing_time"`
	SunOpeningTime  int `form:"sun_opening_time"`
	SunClosingTime  int `form:"sun_closing_time"`
}

type CreateActivityResp struct {
	ActivityId     int64     `json:"activity_id"`
	CreatedAt      time.Time `json:"created_at"`
	ImageSaveFails []string  `json:"failed_images"`
}

type GetActivityReq struct {
	ActivityId int `json:"activity_id"`
}

type GetActivityResp struct {
	Activity    gormModel.Activity `json:"activity"`
	RetrievedAt time.Time          `json:"retrieved_at"`
}

type SearchActivityReq struct {
	SearchText string `json:"search_text"`
	PageSize   int    `json:"page_size"` // assumption: page_size 1 indexed
	PageNumber int    `json:"page_no"`
}

type SearchActivityResp struct {
	Activities   []gormModel.Activity `json:"activities"`
	ResultNumber int64                `json:"result_no"`
}

type UpdateActivityForm struct {
	// Assumption: user token is already validated
	ActivityId  int64    `form:"activity_id"`
	UserId      int64    `form:"user_id"`
	Title       string   `form:"title"`
	Rating      float32  `form:"rating_score"`
	Paid        bool     `form:"paid"`
	Category    []string `form:"category"`
	Description string   `form:"description"`
	Longitude   float32  `form:"longitude"`
	Latitude    float32  `form:"latitude"`
	// Assumption: multiple image file upload at once
	Image []*multipart.FileHeader `form:"image"`

	// fields for opening & closing hours
	MonOpeningTime  int `form:"mon_opening_time"`
	MonClosingTime  int `form:"mon_closing_time"`
	TueOpeningTime  int `form:"tue_opening_time"`
	TueClosingTime  int `form:"tue_closing_time"`
	WedOpeningTime  int `form:"wed_opening_time"`
	WedClosingTime  int `form:"wed_closing_time"`
	ThurOpeningTime int `form:"thur_opening_time"`
	ThurClosingTime int `form:"thur_closing_time"`
	FriOpeningTime  int `form:"fri_opening_time"`
	FriClosingTime  int `form:"fri_closing_time"`
	SatOpeningTime  int `form:"sat_opening_time"`
	SatClosingTime  int `form:"sat_closing_time"`
	SunOpeningTime  int `form:"sun_opening_time"`
	SunClosingTime  int `form:"sun_closing_time"`
}

type UpdateActivityResp struct {
	ActivityId     int64     `json:"activity_id"`
	UpdatedAt      time.Time `json:"updated_at"`
	ImageSaveFails []string  `json:"failed_images"`
}

type InactivateActivityReq struct {
	ActivityId int `json:"activity_id"`
}

type InactivateActivityResp struct {
	ActivityId    int64     `json:"activity_id"`
	InactiveCount int       `json:"inactive_count"`
	InactiveFlag  bool      `json:"inactive_flag"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type DeleteActivityImageReq struct {
	ActivityId int64  `json:"activity_id"`
	UserId     int64  `json:"user_id"`
	ImageName  string `json:"image_name"`
}

type DeleteActivityImageResp struct {
	ActivityId int64     `json:"activity_id"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type CreateReviewReq struct {
	ActivityId int64   `json:"activity_id"`
	UserId     int64   `json:"user_id"`
	Rating     float32 `json:"rating"`
	Review     string  `json:"review"`
}

type CreateReviewResp struct {
	ReviewId      int64     `json:"review_id"`
	CreatedAt     time.Time `json:"created_at"`
	ReviewCounts  int       `json:"review_counts"`
	AverageRating float32   `json:"average_rating"`
}

type UpdateReviewReq struct {
	ReviewId   int     `json:"review_id"`
	ActivityId int64   `json:"activity_id"`
	UserId     int64   `json:"user_id"`
	Delete     bool    `json:"delete"`
	Review     string  `json:"review"`
	Rating     float32 `json:"rating"`
}

type UpdateReviewResp struct {
	ReviewId      int64     `json:"review_id"`
	UpdatedAt     time.Time `json:"Updated_at"`
	ReviewCounts  int       `json:"activity_review_counts"`
	AverageRating float32   `json:"activity_average_rating"`
}
