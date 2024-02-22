package db

import (
	"bibi/biz/dal/mq"
	"gorm.io/gorm"
	"time"
)

//go:generate msgp -tests=false -o=chat_msgp.go -io=false
type Message struct {
	ID        int64		`msg:"id"`
	Uid       int64     `msg:"uid"`
	TargetId  int64     `msg:"target"`
	Content   string    `msg:"content"`
	CreatedAt time.Time `msg:"publish_time"`
	UpdatedAt time.Time `msg:"-"`
	DeletedAt gorm.DeletedAt `sql:"index" msg:"-"`
}

func CreateMessage(midMessage *mq.MiddleMessage) (*Message, error) {
	message := &Message{
		Uid:       midMessage.Uid,
		TargetId:  midMessage.TargetId,
		Content:   midMessage.Content,
		CreatedAt: midMessage.CreatedAt,
	}
	if err := DB.Model(Message{}).Create(message).Error; err != nil {
		return nil, err
	}
	return message, nil
}
