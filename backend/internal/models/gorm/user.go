package gorm

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID        int64          `gorm:"primaryKey;column:id"`
	Username  string         `gorm:"unique;column:username"`
	Email     string         `gorm:"unique;column:email"`
	Interests pq.StringArray `gorm:"type:text[];column:interests"`
	CreatedAt time.Time
}
