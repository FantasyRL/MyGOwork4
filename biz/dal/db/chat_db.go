package db

import (
	"bibi/biz/dal/mq"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        int64
	Uid       int64
	TargetId  int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
}

func CreateMessage(midMessage *mq.MiddleMessage) (*Message, error) {
	message := &Message{
		Uid:      midMessage.Uid,
		TargetId: midMessage.TargetId,
		Content:  midMessage.Content,
	}
	if err := DB.Model(Message{}).Create(message).Error; err != nil {
		return nil, err
	}
	return message, nil
}
