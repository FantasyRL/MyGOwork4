package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             int64  `gorm:"primary_key"`
	Name           string `gorm:"unique"`
	Pwd            string
	FollowCount    int64  `gorm:"default:0"`
	FollowerCount  int64  `gorm:"default:0"`
	Avatar         string `gorm:"default:'https://webstatic.mihoyo.com/bh3/upload/officialsites/201908/ys_1565764084_7084.png'"`
	TotalFavorited int64  `gorm:"default:0"`
	VideoCount     int64  `gorm:"default:0"`
	FavoriteCount  int64  `gorm:"default:0"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
