package follow_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/user"
	"context"
)

type FollowService struct {
	ctx context.Context
}

func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{
		ctx: ctx,
	}
}

func BuildFollowedUsersResp(users []db.User) (usersResp []*user.User) {
	usersResp = make([]*user.User, 0, len(users))
	for _, u := range users {
		usersResp = append(usersResp, &user.User{
			ID:       u.ID,
			Name:     u.UserName,
			Avatar:   u.Avatar,
			IsFollow: true,
		})
	}
	return
}

func BuildFollowerUsersResp(uid int64, users []db.User) (usersResp []*user.User) {
	usersResp = make([]*user.User, 0, len(users))
	for _, u := range users {
		usersResp = append(usersResp, &user.User{
			ID:     u.ID,
			Name:   u.UserName,
			Avatar: u.Avatar,
		})
	}
	return
}
