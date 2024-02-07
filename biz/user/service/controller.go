//MVC--Controller

package service

import (
	"bibi/biz/model/user"
	db2 "bibi/biz/user/dal/db"
	"bibi/pkg/errno"
	"bibi/pkg/utils"
)

func (s *UserService) Register(req *user.RegisterReq) (*db2.User, error) {
	if len(req.Username) < 4 /*||len(req.Password)<8*/ {
		return nil, errno.ParamError
	}

	PwdDigest := utils.SetPassword(req.Password)
	userModel := &db2.User{
		UserName: req.Username,
		Password: PwdDigest,
	}
	return db2.Register(s.ctx, userModel)
}

func (s *UserService) Login(req *user.LoginReq) (*db2.User, error) {
	userModel := &db2.User{
		UserName: req.Username,
		Password: req.Password,
	}
	return db2.Login(s.ctx, userModel)
}

func (s *UserService) Info(id int64) (*db2.User, error) {
	userModel := &db2.User{
		ID: id,
	}
	return db2.QueryUserByID(s.ctx, userModel)
}
