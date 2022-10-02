package gorm

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID        int64          `gorm:"primaryKey;column:id"`
	Username  string         `gorm:"unique;column:username"`
	Email     string         `gorm:"unique;column:email"`
	Password  string         `gorm:"column:password"`
	Interests pq.StringArray `gorm:"type:text[];column:interests"`
	Avatarurl string         `gorm:"column:avatarurl"`
	CreatedAt time.Time
}

type Activity struct {
	ID          int64   `gorm:"primaryKey;column:id"`
	Title       string  `gorm:"column:title"`
	Rating      float32 `gorm:"column:rating"`
	Paid        bool    `gorm:"column:paid"`
	Category    string  `gorm:"column:category"`
	Description string  `gorm:"column:description"`
	Image       string  `gorm:"column:image"`
	OpeningTime int64   `gorm:"column:opening_time"`
	ClosingTime int64   `gorm:"column:closing_time"`
	CreatedAt   time.Time
}

type Itinerary struct {
	ID     int64 `gorm:"primaryKey;column:id"`
	UserID int64 ``
}