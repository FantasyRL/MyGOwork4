package service

import (
	"bibi/biz/model/user"
	"bibi/biz/user/dal/db"
	"context"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func BuildUserResp(_user interface{}) *user.User {
	p, _ := (_user).(*db.User)
	return &user.User{
		ID:     p.ID,
		Name:   p.Name,
		Avatar: p.Avatar,
	}
}
