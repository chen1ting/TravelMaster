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
	UserId     int64  `json:"user_id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	AvatarName string `json:"avatar_file_name"`

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

type GenerateItineraryRequest struct {
	SessionToken        string   `json:"session_token"`
	PreferredCategories []string `json:"preferred_categories"`
	StartTime           int64    `json:"start_time"`
	EndTime             int64    `json:"end_time"`
}

type GenerateItineraryResponse struct {
	GeneratedItinerary *Itinerary `json:"itinerary"`
}

type Itinerary struct {
	Id               int64      `json:"id"`
	Name             string     `json:"name"`
	NumberOfSegments int        `json:"number_of_segments"`
	Segments         []*Segment `json:"segments"`
	StartTime        int64      `json:"start_time"`
	EndTime          int64      `json:"end_time"`
}

type Segment struct {
	StartTime       int64            `json:"start_time"`
	EndTime         int64            `json:"end_time"`
	ActivitySummary *ActivitySummary `json:"activity_summary"`
}

type ActivitySummary struct {
	Id            int64    `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	AverageRating float64  `json:"average_rating"` // to nearest .5 out of 5
	Categories    []string `json:"categories"`
	ImageUrl      string   `json:"image_url"`
}

type GetItineraryRequest struct {
	Id           int64  `json:"id"`
	SessionToken string `json:"session_token"`
}

type GetItineraryResponse struct {
	Itinerary *Itinerary `json:"itinerary"`
}

type GetActivitiesByFilterRequest struct {
	SearchText   string        `json:"search_text"`
	Times        []*TimeFilter `json:"times"`
	SessionToken string        `json:"session_token"`
	PageSize     int64         `json:"page_size"`
	PageNum      int64         `json:"page_num"`
}

type TimeFilter struct {
	Day             int `json:"day"`               // Sunday - Saturday : 0 - 6
	StartTimeOffset int `json:"start_time_offset"` // time offset in hours from 00:00 of that day
	EndTimeOffset   int `json:"end_time_offset"`   // time offset in hours from 00:00 of that day
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

	InactiveCount int        `json:"inactive_count"`
	InactiveFlag  bool       `json:"inactive_flag"`
	ReviewCounts  int        `json:"review_counts"`
	ReviewsList   []*Reviews `json:"review_list"`
	CreatedAt     time.Time  `json:"created_at"`
}

type Reviews struct {
	Id          int64   `json:"id"`
	UserId      int64   `json:"user_id"`
	ActivityId  int64   `json:"activity_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
}

type SearchActivityReq struct {
	SearchText string        `json:"search_text"`
	PageSize   int           `json:"page_size"` // assumption: page_size 1 indexed
	PageNumber int           `json:"page_no"`
	Times      []*TimeFilter `json:"times"`
}

type SearchActivityResp struct {
	NumOfResults int                `json:"num_of_results"`
	Activities   []*ActivitySummary `json:"activities"`
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

type SaveItineraryRequest struct {
	Id           int64      `json:"id"`
	Name         string     `json:"name"`
	SessionToken string     `json:"session_token"`
	Segments     []*Segment `json:"segments"`
}

type SaveItineraryResponse struct {
	Id int64 `json:"id"`
}

type GetItinerariesRequest struct {
	SessionToken string `json:"session_token"`
}

type GetItinerariesResponse struct {
	Itineraries []*Itinerary `json:"itineraries"`
}

type AddReviewReq struct {
	SessionToken string  `json:"session_token"`
	ActivityId   int64   `json:"activity_id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Rating       float32 `json:"rating"`
}

type GetUserInfoReq struct {
	UserId int64 `json:"user_id"`
}

type GetUserInfoResp struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}
