package video_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/user"
	"bibi/biz/model/video"
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

func BuildVideoListResp(videos []db.Video, users []db.User, videoLikeList []int64, isLikeList []int64) []*video.Video {
	var videoListResp []*video.Video
	for i := 0; i < len(videos); i++ {
		cn, _ := time.ParseDuration("8h")
		t := videos[i].CreatedAt.Add(cn)
		videoListResp = append(videoListResp, &video.Video{
			ID:          videos[i].ID,
			Title:       videos[i].Title,
			Author:      BuildAuthorResp(users[i]),
			PlayURL:     videos[i].PlayUrl,
			CoverURL:    videos[i].CoverUrl,
			LikeCount:   videoLikeList[i],
			IsLike:      isLikeList[i],
			PublishTime: t.Format("2006-01-02 15:01:04"),
		})
	}
	return videoListResp
}

func BuildAuthorResp(author db.User) *user.User {
	return &user.User{
		ID:     author.ID,
		Name:   author.UserName,
		Avatar: author.Avatar,
	}
}
