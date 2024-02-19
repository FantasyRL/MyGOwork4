//MVC--Model

package db

import (
	"bibi/pkg/errno"
	"bibi/pkg/utils"
	"context"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             int64
	UserName       string
	Password       string
	FollowCount    int64
	FollowerCount  int64
	Avatar         string
	TotalFavorited int64
	FavoriteCount  int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `sql:"index"`
}

func Register(ctx context.Context, userModel *User) (*User, error) {
	userResp := new(User)
	//WithContext(ctx)是将一个context.Context对象和数据库连接绑定，以实现在数据库操作中使用context.Context上下文传递。
	if err := DB.WithContext(ctx).Where("user_name = ?", userModel.UserName).First(&userResp).Error; err == nil {
		return nil, errno.ExistUserError
	}

	if err := DB.WithContext(ctx).Create(userModel).Error; err != nil {
		return nil, err
	}
	return userModel, nil
}

func Login(ctx context.Context, userModel *User) (*User, error) {
	userResp := new(User)
	if err := DB.WithContext(ctx).Where("user_name = ?", userModel.UserName).
		First(&userResp).Error; err != nil {
		return nil, errno.NotExistUserError
	}

	if utils.CheckPassword(userResp.Password, userModel.Password) == false {
		return nil, errno.PwdError
	}

	return userResp, nil
}

func QueryUserByID(userModel *User) (*User, error) {
	userResp := new(User)
	if err := DB.Model(User{}).Where("id = ?", userModel.ID).First(&userResp).Error; err != nil {
		return nil, err
	}
	return userResp, nil
}

func PutAvatar(ctx context.Context, userModel *User) (*User, error) {
	userResp := new(User)
	if err := DB.Model(User{}).Where("id = ?", userModel.ID).Update("avatar", userModel.Avatar).First(userResp).Error; err != nil {
		return nil, err
	}
	return userResp, nil
}

func QueryUserByIDList(uidList []int64) ([]User, error) {
	userResp := new([]User)
	if err := DB.Model(User{}).Where("id IN ?", uidList).Find(userResp).Error; err != nil {
		return nil, err
	}
	return *userResp, nil
}
