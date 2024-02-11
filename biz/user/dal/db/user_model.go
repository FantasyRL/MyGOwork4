package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             int64
	UserName       string
	Password       string
	FollowCount    int64
	FollowerCount  int64
	Avatar         string
	TotalFavorited int64
	VideoCount     int64
	FavoriteCount  int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
