package db

import "bibi/dao"

func CheckLikeStatus(uid int64, videoId int64, status int64) error {
	var like Like
	return dao.DB.Model(Like{}).Where("uid = ? AND video_id = ? AND status = ?", uid, videoId, status).
		First(&like).Error
}

func IsLikeExist(uid int64, videoId int64) error {
	var like Like
	return dao.DB.Model(Like{}).Where("uid = ? AND video_id = ? ", uid, videoId).
		First(&like).Error
}

func LikeStatusUpdate(uid int64, videoId int64, status int64) error {
	return dao.DB.Model(Like{}).Where("uid = ? AND video_id = ? ", uid, videoId).
		Update("status", status).Error
}

func LikeCreate(uid int64, videoId int64, status int64) error {
	var like = &Like{
		Uid:     uid,
		VideoID: videoId,
		Status:  status,
	}
	return dao.DB.Model(Like{}).Create(like).Error
}
