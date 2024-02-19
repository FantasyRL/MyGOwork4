package follow_service

import (
	"bibi/biz/dal/cache"
	"bibi/biz/dal/db"
	"bibi/biz/model/follow"
	"bibi/pkg/errno"
)

func (s *FollowService) Follow(req *follow.FollowActionReq, uid int64) error {
	//redis
	e, err := cache.IsFollowedCacheExist(s.ctx, req.ObjectUID)
	if err != nil {
		return err
	}
	if e {
		e1, err := cache.IsUserFollowExist(s.ctx, uid, req.ObjectUID)
		if err != nil {
			return err
		}
		if e1 {
			return errno.FollowExistError
		}
		err = cache.AddFollower(s.ctx, uid, req.ObjectUID)
		if err != nil {
			return err
		}
	}

	e1, err := db.IsFollowStatus(uid, req.ObjectUID, 1)
	if err != nil {
		return err
	}
	if e1 {
		return errno.FollowExistError
	}

	e2, err := db.IsFollowExist(uid, req.ObjectUID)
	if err != nil {
		return err
	}
	if !e2 {
		return db.CreateFollow(uid, req.ObjectUID, 1)
	}
	return db.UpdateFollowStatus(uid, req.ObjectUID, 1)
}

func (s *FollowService) UnFollow(req *follow.FollowActionReq, uid int64) error {
	//redis
	e, err := cache.IsFollowedCacheExist(s.ctx, req.ObjectUID)
	if err != nil {
		return err
	}
	if e {
		e1, err := cache.IsUserFollowExist(s.ctx, uid, req.ObjectUID)
		if err != nil {
			return err
		}
		if !e1 {
			return errno.FollowNotExistError
		}
		err = cache.DelFollower(s.ctx, uid, req.ObjectUID)
		if err != nil {
			return err
		}
	}

	//mysql
	e1, err := db.IsFollowStatus(uid, req.ObjectUID, 0)
	if err != nil {
		return err
	}
	if e1 {
		return errno.FollowNotExistError
	}

	e2, err := db.IsFollowExist(uid, req.ObjectUID)
	if err != nil {
		return err
	}
	if !e2 {
		return errno.FollowNotExistError
	}
	return db.UpdateFollowStatus(uid, req.ObjectUID, 0)
}
