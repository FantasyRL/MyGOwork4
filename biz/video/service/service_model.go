package service

import (
	aliyunoss "bibi/oss"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

type VideoService struct {
	ctx    context.Context
	bucket *oss.Bucket
}

func NewVideoService(ctx context.Context) *VideoService {
	bucket, err := aliyunoss.OSSBucketCreate()
	if err != nil {
		log.Fatal(err)
	}
	return &VideoService{ctx: ctx, bucket: bucket}
}
