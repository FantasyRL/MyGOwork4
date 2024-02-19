package follow_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/follow"
)

func (s *FollowService) FollowerList(req *follow.FollowerListReq, uid int64) ([]db.Follow, error) {
	followerList, err := db.FollowerList(uid)
	if err != nil {
		return nil, err
	}
	return followerList, nil
}

func (s *FollowService) FollowingList(req *follow.FollowingListReq, uid int64) ([]db.Follow, error) {
	followedList, err := db.FollowingList(uid)
	if err != nil {
		return nil, err
	}
	return followedList, nil
}

func (s *FollowService) FriendList(req *follow.FriendListReq, uid int64) ([]db.Follow, error) {
	friendList, err := db.FriendList(uid)
	if err != nil {
		return nil, err
	}
	return friendList, nil
}
