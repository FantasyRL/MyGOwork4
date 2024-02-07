//MVC--Model

package db

import (
	"bibi/dao"
	"bibi/pkg/errno"
	"bibi/pkg/utils"
	"context"
)

func Register(ctx context.Context, userModel *User) (*User, error) {
	userResp := new(User)

	//var count int64 = 0
	//dao.DB.WithContext(ctx).Where("name = ?", userModel.UserName).First(&userResp).Count(&count)
	//if count == 1 {
	//	return nil, errno.ExistUserError
	//}
	//WithContext(ctx)是将一个context.Context对象和数据库连接绑定，以实现在数据库操作中使用context.Context上下文传递。
	if err := dao.DB.WithContext(ctx).Where("user_name = ?", userModel.UserName).First(&userResp).Error; err == nil {
		return nil, errno.ExistUserError
	}

	if err := dao.DB.WithContext(ctx).Create(userModel).Error; err != nil {
		return nil, err
	}
	return userModel, nil
}

func Login(ctx context.Context, userModel *User) (*User, error) {
	userResp := new(User)
	if err := dao.DB.WithContext(ctx).Where("user_name = ?", userModel.UserName).
		First(&userResp).Error; err != nil {
		return nil, errno.NotExistUserError
	}

	if utils.CheckPassword(userResp.Password, userModel.Password) == false {
		return nil, errno.PwdError
	}

	return userResp, nil
}

func QueryUserByID(ctx context.Context, userModel *User) (*User, error) {
	userResp := new(User)
	if err := dao.DB.WithContext(ctx).Where("id = ?", userModel.ID).First(&userResp).Error; err != nil {
		return nil, err
	}
	return userResp, nil
}

func PutAvatar(ctx context.Context, userModel *User) (*User, error) {
	userResp := new(User)
	//if err:=dao.DB.WithContext(ctx).Where("id = ?",userModel.ID).Update("avatar",userModel.Avatar).Error;err!=nil{
	//	return nil, err
	//}
	if err := dao.DB.WithContext(ctx).Where("id = ?", userModel.ID).First(userResp).Error; err != nil {
		return nil, err
	}
	userResp.Avatar = userModel.Avatar
	if err := dao.DB.WithContext(ctx).Save(userResp).Error; err != nil {
		return nil, err
	}

	return userModel, nil
}
