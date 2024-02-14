package service

import (
	"context"
)

type LikeService struct {
	ctx context.Context
}

func NewLikeService(ctx context.Context) *LikeService {
	return &LikeService{ctx: ctx}
}
