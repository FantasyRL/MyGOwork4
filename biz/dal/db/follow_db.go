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

func FollowerList(followedId int64) ([]Follow, int64, error) {
	followerList := new([]Follow)
	var count int64
	if err := DB.Model(Follow{}).Where("followed_id = ? AND status = 1", followedId).Find(followerList).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return *followerList, count, nil
}

func FollowingList(uid int64) ([]Follow, int64, error) {
	followedList := new([]Follow)
	var count int64
	if err := DB.Model(Follow{}).Where("uid = ? AND status = 1", uid).Count(&count).Find(followedList).Error; err != nil {
		return nil, 0, err
	}
	return *followedList, count, nil
}

func FriendList(uid int64) ([]Follow, int64, error) {
	friendList1 := new([]Follow)
	friendList2 := new([]Follow)
	var count int64
	if err := DB.Model(Follow{}).Where("uid = ? AND status = 1", uid).Find(friendList1).Error; err != nil {
		return nil, 0, err
	}
	//通过model实现二次查询
	if err := DB.Model(&friendList1).Where("followed_id = ? AND status = 1", uid).Find(friendList2).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return *friendList2, count, nil
}

func IsFollow(uid int64, followedId int64) (bool, error) {
	follow := new(Follow)
	err := DB.Model(Follow{}).Where("uid = ? AND followed_id = ?", uid, followedId).First(follow).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
