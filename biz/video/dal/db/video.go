package db

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	//db.User   `gorm:"ForeignKey:Uid"`
	ID        int64 `gorm:"primary_key"`
	UserName  string
	Uid       int64
	Title     string
	PlayUrl   string
	CoverUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
}
