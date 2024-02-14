package user_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/user"
	aliyunoss "bibi/oss"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

type UserService struct {
	ctx context.Context
}

type AvatarService struct {
	ctx    context.Context
	bucket *oss.Bucket
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func NewAvatarService(ctx context.Context) *AvatarService {
	bucket, err := aliyunoss.OSSBucketCreate()
	if err != nil {
		log.Fatal(err)
	}
	return &AvatarService{ctx: ctx, bucket: bucket}
}

func BuildUserResp(_user interface{}) *user.User {
	//这里使用了一个及其抽象的断言
	p, _ := (_user).(*db.User)
	return &user.User{
		ID:     p.ID,
		Name:   p.UserName,
		Avatar: p.Avatar,
	}
}
