package service

import (
	"bibi/biz/model/video"
	"bibi/biz/video/dal/db"
)

func (s *VideoService) SearchVideo(req *video.SearchVideoReq) ([]db.Video, int64, error) {
	return db.SearchVideo(s.ctx, int(req.PageNum), req.Param)
}
