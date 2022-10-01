package gorm

import (
	"github.com/lib/pq"
	"time"
)

type User struct {
	ID        int64          `gorm:"primaryKey;column:id"`
	Username  string         `gorm:"unique;column:username"`
	Email     string         `gorm:"unique;column:email"`
	Password  string         `gorm:"column:password"`
	Interests pq.StringArray `gorm:"type:text[];column:interests"`
	CreatedAt time.Time
}

type Activity struct {
	// Basic information on the activity
	ID            int64          `gorm:"primaryKey;column:id"`
	UserID        int64          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id"`
	Title         string         `gorm:"unique;not null;type:varchar(100);default:null;column:title"`
	AverageRating float32        `gorm:"column:rating"`
	Paid          bool           `gorm:"column:paid"`
	Category      pq.StringArray `gorm:"type:text[];column:category"`
	Description   string         `gorm:"column:description"`
	Longitude     float32        `gorm:"gorm:longitude; default:-180.1"`
	Latitude      float32        `gorm:"gorm:latitude; default:-90.1"`
	ImageURL      string         `gorm:"column:image"`
	OpeningTimes  pq.Int32Array  `gorm:"type:int[];column:opening_times"`
	// 0:Mon Opening, 1: Tue Opening  2: Wed Opening ... 7: Mon Closing
	// fields for opening & closing hours
	/*
		MonOpeningTime  int `gorm:"column:mon_opening_time; default:-1"`
		MonClosingTime  int `gorm:"column:mon_closing_time; default:-1"`
		TueOpeningTime  int `gorm:"column:tue_opening_time; default:-1"`
		TueClosingTime  int `gorm:"column:tue_closing_time; default:-1"`
		WedOpeningTime  int `gorm:"column:wed_opening_time; default:-1"`
		WedClosingTime  int `gorm:"column:wed_closing_time; default:-1"`
		ThurOpeningTime int `gorm:"column:thur_opening_time; default:-1"`
		ThurClosingTime int `gorm:"column:thur_closing_time; default:-1"`
		FriOpeningTime  int `gorm:"column:fri_opening_time; default:-1"`
		FriClosingTime  int `gorm:"column:fri_closing_time; default:-1"`
		SatOpeningTime  int `gorm:"column:sat_opening_time; default:-1"`
		SatClosingTime  int `gorm:"column:sat_closing_time; default:-1"`
		SunOpeningTime  int `gorm:"column:sun_opening_time; default:-1"`
		SunClosingTime  int `gorm:"column:sun_closing_time; default:-1"`
	*/

	// System fields
	InactiveCount int    `gorm:"column:inactive_count"`
	InactiveFlag  bool   `gorm:"column:inactive_flag"`
	ReviewCounts  int    `gorm:"column:review_counts"`
	ReviewIds     string `gorm:"column:review_ids"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Itinerary struct {
	ID     int64 `gorm:"primaryKey;column:id"`
	UserID int64 ``
}
