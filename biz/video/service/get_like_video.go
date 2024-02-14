package service

import (
	"bibi/biz/model/video"
	"bibi/biz/video/dal/db"
)

var test video.Video

//todo:rpc
func (s VideoService) GetLikeVideoList(videoIdList []int64) ([]db.Video, error /*,likeList []int64, isLikeList []int64*/) {
	return db.GetVideoByIdList(videoIdList)
}
