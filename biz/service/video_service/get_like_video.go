package video_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/video"
)

var test video.Video

//todo:rpc
func (s *VideoService) GetLikeVideoList(videoIdList []int64) ([]db.Video, error /*,likeList []int64, isLikeList []int64*/) {
	return db.GetVideoByIdList(videoIdList)
}
