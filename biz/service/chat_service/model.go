package chat_service

import (
	"context"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService { return &MessageService{ctx: ctx} }

type SendMsg struct {
	Type    int64  `json:"type"`
	Content string `json:"content"`
}
