package db

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	//db.User   `gorm:"ForeignKey:Uid"`
	ID        int64  `gorm:"primary_key"`
	UserName  string `gorm:"not null"`
	Tid       int64  `gorm:"not null"`
	Uid       int64  `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	View      int64  `gorm:"default:'0'"`
	Status    int64  `gorm:"default:'0'"` //0正常 1封禁
	Video     string `gorm:"type:longblob"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
}
