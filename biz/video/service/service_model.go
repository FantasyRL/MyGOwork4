package service

import (
	"bibi/biz/model/video"
	"bibi/biz/video/dal/db"
	aliyunoss "bibi/oss"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"time"
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
func BuildVideoResp(v *db.Video) *video.Video {
	cn, _ := time.ParseDuration("8h")
	t := v.CreatedAt.Add(cn)
	return &video.Video{
		ID:          v.ID,
		UID:         v.Uid,
		Title:       v.Title,
		PlayURL:     v.PlayUrl,
		CoverURL:    v.CoverUrl,
		PublishTime: t.Format("2006-01-02 15:01:04"),
	}
}
func BuildListVideoResp(list []db.Video) (videos []*video.Video) {
	for _, v := range list {
		videos = append(videos, BuildVideoResp(&v))
	}
	return
}
