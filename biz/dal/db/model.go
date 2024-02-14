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
	FavoriteCount  int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `sql:"index"`
}

type Video struct {
	ID        int64 `gorm:"primary_key"`
	Uid       int64
	Title     string
	PlayUrl   string
	CoverUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
}

type Like struct {
	ID        int64
	Uid       int64
	VideoID   int64
	Status    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
}
