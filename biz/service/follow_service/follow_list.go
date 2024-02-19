package follow_service

import (
	"bibi/biz/dal/cache"
	"bibi/biz/dal/db"
	"bibi/biz/model/follow"
)

func (s *FollowService) FollowerList(req *follow.FollowerListReq, uid int64) ([]db.Follow, int64, error) {
	//redis
	followerList, err := cache.GetFollower(s.ctx, uid)
	if err != nil {
		return nil, 0, err
	}
	e1, count, err := cache.GetFollowerCount(s.ctx, uid)
	if err != nil {
		return nil, 0, err
	}

	if len(followerList) != 0 && e1 {
		return followerList, count, nil
	}

	//mysql
	followerList, count, err = db.FollowerList(uid)
	if err != nil {
		return nil, 0, err
	}
	if err = cache.SetFollowerList(s.ctx, uid, followerList); err != nil {
		return nil, 0, err
	}

	if err = cache.SetFollowerCount(s.ctx, uid, count); err != nil {
		return nil, 0, err
	}

	return followerList, count, nil

}

func (s *FollowService) FollowingList(req *follow.FollowingListReq, uid int64) ([]db.Follow, int64, error) {
	followedList, count, err := db.FollowingList(uid)
	if err != nil {
		return nil, 0, err
	}
	if err = cache.SetFollowingList(s.ctx, uid, followedList); err != nil {
		return nil, 0, err
	}
	if err = cache.SetFollowingCount(s.ctx, uid, count); err != nil {
		return nil, 0, err
	}

	return followedList, count, nil
}

func (s *FollowService) FriendList(req *follow.FriendListReq, uid int64) ([]db.Follow, int64, error) {
	friendList, count, err := db.FriendList(uid)
	if err != nil {
		return nil, 0, err
	}
	return friendList, count, nil
}
