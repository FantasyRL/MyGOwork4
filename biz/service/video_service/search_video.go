package video_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/video"
)

func (s *VideoService) SearchVideo(req *video.SearchVideoReq) ([]db.Video, int64, error) {
	return db.SearchVideo(s.ctx, int(req.PageNum), req.Param)
}
