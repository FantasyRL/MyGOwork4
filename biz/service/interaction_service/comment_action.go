package interaction_service

import (
	"bibi/biz/dal/cache"
	"bibi/biz/dal/db"
	"bibi/biz/model/interaction"
	"bibi/biz/model/user"
	"bibi/pkg/errno"
	"golang.org/x/sync/errgroup"
)

func (s *InteractionService) CommentCreate(req *interaction.CommentCreateReq, uid int64) (*db.Comment, error) {
	var eg errgroup.Group
	var err error
	var exist = false
	var comment *db.Comment
	eg.Go(func() error {
		var commentModel = &interaction.Comment{
			VideoID: req.VideoID,
			Content: req.Content,
			User: &user.User{
				ID: uid,
			},
		}
		//若内容完全重复，则删除最早发的那个(其实是懒得再开一个接口了)
		comment, err = db.CreateComment(commentModel)
		if err != nil {
			return err
		}
		return cache.AddVideoComment(s.ctx, comment)
	})

	eg.Go(func() error {
		exist, _, err = cache.GetVideoCommentCount(s.ctx, req.VideoID)
		if err != nil {
			return err
		}
		if exist {
			return cache.IncrVideoCommentCount(s.ctx, req.VideoID)
		}
		return nil
	})

	if err = eg.Wait(); err != nil {
		return nil, err
	}
	if !exist {
		count, err := db.GetCommentCount(req.VideoID)
		if err != nil {
			return nil, err
		}
		err = cache.SetVideoCommentCount(s.ctx, req.VideoID, count)
		if err != nil {
			return nil, err
		}
	}
	return comment, nil
}

func (s *InteractionService) CommentDelete(req *interaction.CommentDeleteReq, uid int64) error {
	var eg errgroup.Group
	var commentModel = &interaction.Comment{
		ID:      req.CommentID,
		VideoID: req.VideoID,
		User: &user.User{
			ID: uid,
		},
	}
	exist, err := db.IsCommentExist(commentModel)
	if err != nil {
		return err
	}
	if !exist {
		return errno.CommentIsNotExistError
	}
	eg.Go(func() error {

		comment, err := db.DeleteComment(commentModel)
		if err != nil {
			return err
		}
		return cache.DelVideoComment(s.ctx, comment)
	})

	eg.Go(func() error {
		exist, _, err := cache.GetVideoCommentCount(s.ctx, req.VideoID)
		if err != nil {
			return err
		}
		if exist {
			return cache.DecrVideoCommentCount(s.ctx, req.VideoID)
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
