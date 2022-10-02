package models

import (
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
	Interests      []string              `form:"interests"`
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

type CreateActivityForm struct {
	// Assumption: user token is already validated
	UserId      int64    `form:"user_id"`
	Title       string   `form:"title"`
	Rating      float32  `form:"rating_score"`
	Paid        bool     `form:"paid"`
	Category    []string `form:"category"`
	Description string   `form:"description"`
	Longitude   float32  `form:"longitude"`
	Latitude    float32  `form:"latitude"`
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
	ActivityId  int64    `json:"activity_id"`
	Title       string   `json:"title"`
	Rating      float32  `json:"rating_score"`
	Paid        bool     `json:"paid"`
	Category    []string `json:"category"`
	Description string   `json:"description"`
	Longitude   float32  `json:"longitude"`
	Latitude    float32  `json:"latitude"`
	ImageNames  []string `json:"image_names"`

	// fields for opening & closing hours
	MonOpeningTime  int `json:"mon_opening_time"`
	MonClosingTime  int `json:"mon_closing_time"`
	TueOpeningTime  int `json:"tue_opening_time"`
	TueClosingTime  int `json:"tue_closing_time"`
	WedOpeningTime  int `json:"wed_opening_time"`
	WedClosingTime  int `json:"wed_closing_time"`
	ThurOpeningTime int `json:"thur_opening_time"`
	ThurClosingTime int `json:"thur_closing_time"`
	FriOpeningTime  int `json:"fri_opening_time"`
	FriClosingTime  int `json:"fri_closing_time"`
	SatOpeningTime  int `json:"sat_opening_time"`
	SatClosingTime  int `json:"sat_closing_time"`
	SunOpeningTime  int `json:"sun_opening_time"`
	SunClosingTime  int `json:"sun_closing_time"`

	InactiveCount int       `json:"inactive_count"`
	InactiveFlag  bool      `json:"inactive_flag"`
	ReviewCounts  int       `json:"review_counts"`
	ReviewList    string    `json:"review_list"`
	CreatedAt     time.Time `json:"created_at"`
}

type SearchActivityReq struct {
	SearchText string `json:"search_text"`
	PageSize   int    `json:"page_size"` // assumption: page_size 1 indexed
	PageNumber int    `json:"page_no"`
}

type SearchActivityResp struct {
	Activities   string `json:"activities"`
	ResultNumber int64  `json:"result_no"`
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
