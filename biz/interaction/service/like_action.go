package service

import (
	"bibi/biz/interaction/dal/cache"
	"bibi/biz/interaction/dal/db"
	"bibi/biz/model/interaction"
	"bibi/pkg/errno"
	"errors"
	"gorm.io/gorm"
)

//todo:isVideoExist;isAuthor(uid:video_id:countSuffix)
func (s *LikeService) Like(req *interaction.LikeActionReq, uid int64) error {
	//video是否存在，那不是还是在和数据库交互
	//if err=dbVideo.CheckVideoExistById(req.VideoID);err!=nil{
	//	return
	//}
	//点赞是否存在于redis
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoID, uid)
	if err != nil {
		return err
	}
	if exist {
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
	}

	//检查点赞是否重复
	if err = db.CheckLikeStatus(uid, req.VideoID, 1); err == nil {
		return errno.LikeExistError
	}

	//检查点赞是否存在
	err = db.IsLikeExist(uid, req.VideoID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//创建点赞
		return db.LikeCreate(uid, req.VideoID, 1)
	}
	return db.LikeStatusUpdate(uid, req.VideoID, 1)

}

func (s *LikeService) DisLike(req *interaction.LikeActionReq, id int64) error {
	return nil

}
