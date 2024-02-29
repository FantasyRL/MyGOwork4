package chat_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/chat"
	"context"
	"time"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService { return &MessageService{ctx: ctx} }

type SendMsg struct {
	Type    int64  `json:"type"`
	Content string `json:"content"`
}

func BuildMessageResp(msgList []db.Message) []*chat.Message {
	var msgs []*chat.Message
	for _, msg := range msgList {
		msgs = append(msgs, &chat.Message{
			ID:         msg.ID,
			TargetID:   msg.TargetId,
			FromID:     msg.Uid,
			Content:    msg.Content,
			CreateTime: msg.CreatedAt.Format(time.RFC3339),
		})
	}
	return msgs
}
