package pack

import (
	"bibi/biz/model/user"
	"bibi/biz/model/video"
	db2 "bibi/biz/user/dal/db"
	"bibi/biz/video/dal/db"
	"time"
)

func BuildVideoListResp(videos []db.Video, users []db2.User, videoLikeList []int64, isLikeList []int64) []*video.Video {
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

func BuildAuthorResp(author db2.User) *user.User {
	return &user.User{
		ID:     author.ID,
		Name:   author.UserName,
		Avatar: author.Avatar,
	}
}
