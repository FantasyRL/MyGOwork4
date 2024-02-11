package db

import (
	"bibi/dao"
	"bibi/pkg/conf"
	"bibi/pkg/errno"
	"context"
)

func CreateVideo(ctx context.Context, video *Video) (*Video, error) {
	if err := dao.DB.WithContext(ctx).Create(video).Error; err != nil {
		return nil, err
	}
	return video, nil
}

func GetVideoCountByID(uid int64) (count int64, err error) {
	if err = dao.DB.Model(Video{}).Where("uid = ?", uid).Count(&count).Error; err != nil {
		return 114514, err
	}
	return
}

func ListVideosByID(ctx context.Context, pageNum int, uid int64) ([]Video, int64, error) {
	videos := new([]Video)
	var count int64 = 0
	if err := dao.DB.Model(&Video{}).Where("uid = ?", uid).Count(&count).Order("created_at DESC").
		Limit(conf.PageSize).Offset((pageNum - 1) * conf.PageSize).Find(videos).
		Error; err != nil {
		return nil, 114514, err
	}
	return *videos, count, nil
}

func SearchVideo(ctx context.Context, pageNum int, param string) ([]Video, int64, error) {
	videos := new([]Video)
	var count int64 = 0
	if err := dao.DB.Model(&Video{}).
		Where("title LIKE ? ", "%"+param+"%").
		Count(&count).Limit(conf.PageSize).Offset((pageNum - 1) * conf.PageSize).Find(&videos).
		Error; err != nil {
		return nil, 114514, err
	}
	return *videos, count, nil
}

func CheckVideoExistById(videoId int64) error {
	video := new(Video)
	if err := dao.DB.Model(Video{}).Where("id = ?", videoId).Find(video).Error; err != nil {
		return err
	}
	if video != (&Video{}) {
		return errno.VideoNotExistError
	}
	return nil
}
