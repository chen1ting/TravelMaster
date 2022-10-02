package gorm

import (
	"github.com/lib/pq"
	"time"
)

type User struct {
	ID         int64          `gorm:"primaryKey;column:id"`
	Username   string         `gorm:"unique;column:username"`
	Email      string         `gorm:"unique;column:email"`
	Password   string         `gorm:"column:password"`
	Interests  pq.StringArray `gorm:"type:text[];column:interests"`
	AboutMe    string         `gorm:"column:about_me"`
	AvatarName string         `gorm:"column:avatar_name"`
	Activities []Activity
	Reviews    []Review
	CreatedAt  time.Time
}

type Activity struct {
	// Basic information on the activity
	ID              int64 `gorm:"primaryKey;column:id"`
	UserID          int64
	Title           string         `gorm:"unique;not null;type:varchar(100);default:null;column:title"`
	AuthorRating    float32        `gorm:"column:author_rating"`
	AvgReviewRating float32        `gorm:"column:average_rating"`
	Paid            bool           `gorm:"column:paid"`
	Category        pq.StringArray `gorm:"type:text[];column:category"`
	Description     string         `gorm:"column:description"`
	Longitude       float32        `gorm:"gorm:longitude; default:-180.1"`
	Latitude        float32        `gorm:"gorm:latitude; default:-90.1"`
	ImageNames      pq.StringArray `gorm:"type:text[];column:image_name"`
	OpeningTimes    pq.Int32Array  `gorm:"type:int[];column:opening_times"`

	// System fields
	InactiveCount int  `gorm:"column:inactive_count; default: 0"`
	InactiveFlag  bool `gorm:"column:inactive_flag; default:false"`
	ReviewCounts  int  `gorm:"column:review_counts; default:0"`
	Reviews       []Review
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Itinerary struct {
	ID     int64 `gorm:"primaryKey;column:id"`
	UserID int64 ``
}

type Review struct {
	ID         int64 `gorm:"primaryKey;column:id"`
	UserId     int
	User       User
	ActivityId int
	Activity   Activity
	Review     string  `gorm:"column:review"`
	Rating     float32 `gorm:"column:rating"`
}
