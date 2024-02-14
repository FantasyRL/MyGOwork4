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

func GetVideoByUid(uid int64) ([]int64, error) {
	likes := new([]Like)
	if err := dao.DB.Model(Like{}).Where("uid = ? AND status = ?", uid, 1).Find(likes).Error; err != nil {
		return nil, err
	}
	var videoIdList []int64
	for _, id := range *likes {
		videoIdList = append(videoIdList, id.VideoID)
	}
	return videoIdList, nil
}

func GetVideoLikeCount(videoId int64) (count int64, err error) {
	if err = dao.DB.Model(Like{}).Where("video_id = ?", videoId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
