package models

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

type SignupReq struct {
	Username       string   `json:"username"`
	HashedPassword string   `json:"hashed_password"`
	Email          string   `json:"email"`
	// Interests      []string `json:"interests"` update later, not as part of initial sign up req
}

type SignupResp struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	SessionToken string `json:"session_token"`
}

type ValidateTokenReq struct {
	SessionToken string `json:"session_token"`
}

type ValidateTokenResp struct {
	Valid bool `json:"valid"`
	UserId int64 `json:"user_id"`
}


type GenerateItineraryRequest struct {
	SessionToken string `json:"session_token"`
	PreferredCategories []string `json:"preferred_categories"`
	StartTime int64 `json:"start_time"`
	EndTime int64 `json:"end_time"`
}

type GenerateItineraryResponse struct {
	GeneratedItinerary *Itinerary `json:"itinerary"`
}

type Itinerary struct {
	Id int64 `json:"id"`
	NumberOfSegments int `json:"number_of_segments"`
	Segments []*Segment `json:"segments"`
	StartTime int64 `json:"start_time"`
	EndTime int64 `json:"end_time"`
}

type Segment struct {
	StartTime int `json:"start_time"` // every time bin is 30 minutes, time bin 0 = 00:00, time bin 1 = 00:29, etc
	EndTime int `json:"end_time"` // in total there will be 48 time bins from 00:00 to 23:59
	ActivitySummary ActivitySummary `json:"activity_summary"`
}

type ActivitySummary struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	AverageRating float64 `json:"average_rating"` // to nearest .5 out of 5
	Categories []string `json:"categories"`
	ImageUrl string `json:"image_url"`
}

type GetItineraryRequest struct {
	Id int64 `json:"id"`
	SessionToken string `json:"session_token"`
}

type GetItineraryResponse struct {
	Itinerary *Itinerary `json:"itinerary"`
}

type GetActivitiesByFilterRequest struct {
	SearchText string `json:"search_text"`
	Times []*TimeFilter `json:"times"`
	SessionToken string `json:"session_token"`
	PageSize int64 `json:"page_size"`
	PageNum int64 `json:"page_num"`
}

type TimeFilter struct {
	Day int `json:"day"`// Sunday - Saturday : 0 - 6
	StartTimeOffset int `json:"start_time_offset"` // time offset in hours from 00:00 of that day
	EndTimeOffset int `json:"end_time_offset"` // time offset in hours from 00:00 of that day
}

type GetActivitiesByFilterResponse struct {
	NumOfResults int `json:"num_of_results"`
	Activities []*ActivitySummary `json:"activities"`
}
