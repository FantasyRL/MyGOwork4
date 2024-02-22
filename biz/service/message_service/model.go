package message_service

import "context"

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService { return &MessageService{ctx: ctx} }
