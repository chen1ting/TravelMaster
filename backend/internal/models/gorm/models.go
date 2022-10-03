package gorm

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type User struct {
	ID         int64          `gorm:"primaryKey;column:id"`
	Username   string         `gorm:"unique;column:username"`
	Email      string         `gorm:"unique;column:email"`
	Password   string         `gorm:"column:password"`
	Interests  pq.StringArray `gorm:"type:text[];column:interests"`
	AvatarName string         `gorm:"column:avatar_name"`
	CreatedAt  time.Time
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
	ImageNames    pq.StringArray `gorm:"type:text[];column:image_name"`
	OpeningTimes  pq.Int32Array  `gorm:"type:int[];column:opening_times"`

	// System fields
	InactiveCount int           `gorm:"column:inactive_count"`
	InactiveFlag  bool          `gorm:"column:inactive_flag"`
	ReviewCounts  int           `gorm:"column:review_counts"`
	ReviewIds     pq.Int64Array `gorm:"type:int[];column:review_ids"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserReports   []User `gorm:"many2many:user_reports"`
}

type Itinerary struct {
	ID               int64          `gorm:"primaryKey;column:id"`
	Name             string         `gorm:"column:name"`
	OwnedByUserId    int64          `gorm:"column:owned_by_user_id"` // TODO: Foreign key to User id
	Segments         datatypes.JSON `gorm:"type:jsonb;column:segments"`
	StartTime        int64          `gorm:"column:start_time"`
	EndTime          int64          `gorm:"column:end_time"`
	NumberOfSegments int            `gorm:"column:num_of_segments"`
}

type Review struct {
	ID          int64   `gorm:"primaryKey;column:id"`
	Title       string  `gorm:"column:title"`
	Description string  `gorm:"column:description"`
	UserId      int64   `gorm:"uniqueIndex:unique_review"` // TODO: Foreign key to User id
	ActivityId  int64   `gorm:"uniqueIndex:unique_review"` // TODO: Foreign key to Activity id
	Rating      float32 `gorm:"column:rating"`
}

type ReportHistory struct {
	UserId     int64  `gorm:"primaryKey;column:user_id"`
	ActivityId int64  `gorm:"primaryKey;column:activity_id"`
	Reason     string `gorm:"column:reason"` //for feedback purposes
	CreatedAt  time.Time
}
