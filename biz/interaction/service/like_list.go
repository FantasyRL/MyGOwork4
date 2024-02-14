package service

import (
	"bibi/biz/interaction/dal/cache"
	"bibi/biz/interaction/dal/db"
	"bibi/biz/model/interaction"
	"errors"
	"gorm.io/gorm"
)

func (s *LikeService) LikeVideoList(req *interaction.LikeListReq, uid int64) ([]int64, error) {
	//缓存未过期
	videoIdList, err := cache.GetUserLikeVideos(s.ctx, uid)
	if err != nil {
		return nil, err
	}
	if len(videoIdList) != 0 {
		return videoIdList, nil
	}

	//缓存过期
	videoIdList, err = db.GetVideoByUid(uid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	//将mysql数据存入redis缓存
	err = cache.AddLikeVideoList(s.ctx, videoIdList, uid)
	if err != nil {
		return nil, err
	}
	return videoIdList, nil
}
