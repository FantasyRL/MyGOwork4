package db

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Follow struct {
	ID         int64 `gorm:"primary_key"`
	Uid        int64 `gorm:"uid"`
	FollowedId int64 `gorm:"followed_id"`
	Status     int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `sql:"index"`
}

func IsFollowExist(uid int64, followerId int64) (bool, error) {
	follow := new(Follow)
	err := DB.Model(Follow{}).Where("uid = ? AND followed_id = ? ", uid, followerId).First(follow).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func IsFollowStatus(uid int64, followerId int64, status int64) (bool, error) {
	follow := new(Follow)
	err := DB.Model(Follow{}).Where("uid = ? AND followed_id = ? AND status = ? ", uid, followerId, status).First(follow).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateFollow(uid int64, followerId int64, status int64) error {
	followModel := new(Follow)
	followModel = &Follow{
		Uid:        uid,
		FollowedId: followerId,
		Status:     status,
	}
	return DB.Model(Follow{}).Create(followModel).Error
}

func UpdateFollowStatus(uid int64, followerId int64, status int64) error {
	return DB.Model(Follow{}).Where("uid = ? AND followed_id = ? ",
		uid, followerId).Update("status", status).Error
}

func FollowerList(followed_id int64) ([]Follow, error) {
	followerList := new([]Follow)
	if err := DB.Model(Follow{}).Where("followed_id = ? AND status = 1", followed_id).Find(followerList).Error; err != nil {
		return nil, err
	}
	return *followerList, nil
}

func FollowingList(uid int64) ([]Follow, error) {
	followedList := new([]Follow)
	if err := DB.Model(Follow{}).Where("uid = ? AND status = 1", uid).Find(followedList).Error; err != nil {
		return nil, err
	}
	return *followedList, nil
}

func FriendList(uid int64) ([]Follow, error) {
	friendList := new([]Follow)
	if err := DB.Model(Follow{}).Where("uid = ? AND status = 1", uid).
		Where("followed_id = ? AND status = 1", uid).Find(friendList).Error; err != nil {
		return nil, err
	}
	return *friendList, nil
}

func IsFollow(uid int64, followed_id int64) (bool, error) {
	follow := new(Follow)
	err := DB.Model(Follow{}).Where("uid = ? AND followed_id = ?", uid, followed_id).First(follow).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
