package interaction_service

import (
	"bibi/biz/dal/cache"
	"bibi/biz/dal/db"
	"bibi/biz/model/interaction"
	"bibi/pkg/errno"
	"errors"
	"gorm.io/gorm"
)

//todo:isVideoExist;isAuthor(uid:video_id:countSuffix)
func (s *InteractionService) Like(req *interaction.LikeActionReq, uid int64) error {

	//用户数据是否存在于redis
	exist, err := cache.IsUserLikeCacheExist(s.ctx, uid)
	if err != nil {
		return err
	}
	if !exist {
		videoIdList, err := db.GetVideoByUid(uid)
		if err != nil {
			return err
		}
		err = cache.AddLikeVideoList(s.ctx, videoIdList, uid)
		if err != nil {
			return err
		}
	}

	//该点赞是否存在
	exist2, err := cache.IsVideoLikeExist(s.ctx, req.VideoID, uid)
	if err != nil {
		return err
	}
	if exist2 {
		return errno.LikeExistError
	}

	//视频点赞量redis是否过期,若过期则直接存入mysql，未过期则同步视频点赞量
	ok, _, err := cache.GetVideoLikeCount(s.ctx, req.VideoID)
	if err != nil {
		return err
	}
	//video存在
	if ok {
		//向redis添加用户点赞与视频点赞量
		if err := cache.AddVideoLikeCount(s.ctx, req.VideoID, uid); err != nil {
			return err
		}

	} else {
		//只添加用户点赞
		if err := cache.AddUserLikeVideo(s.ctx, req.VideoID, uid); err != nil {
			return err
		}
	}

	//检查点赞条目是否存在，存在则更新，不存在则创建
	err = db.IsLikeExist(uid, req.VideoID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//创建点赞
		return db.LikeCreate(uid, req.VideoID, 1)
	}
	return db.LikeStatusUpdate(uid, req.VideoID, 1)
}

func (s *InteractionService) DisLike(req *interaction.LikeActionReq, uid int64) error {
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoID, uid)
	if err != nil {
		return err
	}
	if !exist {
		err = db.IsLikeExist(uid, req.VideoID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.LikeNotExistError
		}
		if err != nil {
			return err
		}

		if err = db.CheckLikeStatus(uid, req.VideoID, 0); err == nil {
			return errno.LikeNotExistError
		}
	}
	if exist {
		if err = cache.DelVideoLikeCount(s.ctx, req.VideoID, uid); err != nil {
			return err
		}
	}
	return db.LikeStatusUpdate(uid, req.VideoID, 0)
}
