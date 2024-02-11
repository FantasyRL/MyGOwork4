package service

import (
	"bibi/biz/model/video"
	"bibi/biz/video/dal/db"
)

func (s *VideoService) ListVideo(req *video.ListUserVideoReq, uid int64) ([]db.Video, int64, error) {
	return db.ListVideosByID(s.ctx, int(req.PageNum), uid)
}
