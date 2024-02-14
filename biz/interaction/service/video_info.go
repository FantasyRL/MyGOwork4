package service

import (
	"bibi/biz/interaction/dal/cache"
	"bibi/biz/interaction/dal/db"
	"errors"
	"gorm.io/gorm"
)

//todo:commentCount

func (s *LikeService) GetVideoLikeById(videoId int64) (int64, error) {
	//redis
	_, likeCount, err := cache.GetVideoLikeCount(s.ctx, videoId)
	if err != nil {
		return 0, err
	}
	if likeCount != 0 {
		return likeCount, nil
	}
	//db
	likeCount, err = db.GetVideoLikeCount(videoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	//存入redis
	if err = cache.SetVideoLikeCounts(s.ctx, videoId, likeCount); err != nil {
		return 0, err
	}
	return likeCount, nil
}
