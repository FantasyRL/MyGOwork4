package video_service

import (
	"bibi/biz/dal/cache"
	"bibi/biz/dal/db"
	"bibi/biz/model/video"
)

func (s *VideoService) HotVideo(req *video.HotVideoReq) ([]db.Video, error) {
	videoIdList, err := cache.ListHotVideo(s.ctx)
	if err != nil {
		return nil, err
	}
	videoList, err := db.GetVideoByIdList(videoIdList)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
