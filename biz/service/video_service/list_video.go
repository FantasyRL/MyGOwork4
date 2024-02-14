package video_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/video"
)

func (s *VideoService) ListVideo(req *video.ListUserVideoReq, uid int64) ([]db.Video, int64, error) {
	return db.ListVideosByID(s.ctx, int(req.PageNum), uid)
}
