package db

import (
	"bibi/dao"
	"context"
)

func CreateVideo(ctx context.Context, video *Video) (*Video, error) {
	if err := dao.DB.WithContext(ctx).Create(video).Error; err != nil {
		return nil, err
	}
	return video, nil
}
