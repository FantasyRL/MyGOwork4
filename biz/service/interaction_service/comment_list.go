package interaction_service

import (
	"bibi/biz/dal/cache"
	"bibi/biz/dal/db"
	"bibi/biz/model/interaction"
	"errors"
	"gorm.io/gorm"
)

func (s *InteractionService) CommentList(req *interaction.CommentListReq) ([]db.Comment, int64, error) {
	commentCache, err := cache.GetVideoComments(s.ctx, req.VideoID)
	if err != nil {
		return nil, 0, err
	}
	exist, countCache, err := cache.GetVideoCommentCount(s.ctx, req.VideoID)
	if err != nil {
		return nil, 0, err
	}
	if exist && len(commentCache) != 0 {
		return commentCache, countCache, nil
	}
	if !exist && len(commentCache) != 0 {
		count, err := db.GetCommentCount(req.VideoID)
		if err != nil {
			return nil, 0, err
		}
		return commentCache, count, nil
	}
	comments, count, err := db.GetCommentsByVideoID(req.VideoID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	}
	if err != nil {
		return nil, 0, err
	}
	//设置缓存
	if err := cache.SetVideoComments(s.ctx, comments, req.VideoID); err != nil {
		return nil, 0, err
	}
	return comments, count, nil
}
