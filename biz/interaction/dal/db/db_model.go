package db

import (
	"gorm.io/gorm"
	"time"
)

type Like struct {
	ID        int64
	Uid       int64
	VideoID   int64
	Status    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
