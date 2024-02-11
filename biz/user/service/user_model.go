package service

import (
	"bibi/biz/model/user"
	"bibi/biz/user/dal/db"
	aliyunoss "bibi/oss"
	"bibi/pkg/errno"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"log"
	"mime/multipart"
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
	p, _ := (_user).(*db.User)
	return &user.User{
		ID:     p.ID,
		Name:   p.UserName,
		Avatar: p.Avatar,
	}
}

func FileToByte(file *multipart.FileHeader) ([]byte, error) {
	fileContent, err := file.Open()
	if err != nil {
		return nil, errno.ParamError
	}
	return io.ReadAll(fileContent)
}
