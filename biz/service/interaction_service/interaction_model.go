package interaction_service

import (
	"context"
)

type InteractionService struct {
	ctx context.Context
}

func NewInteractionService(ctx context.Context) *InteractionService {
	return &InteractionService{ctx: ctx}
}
