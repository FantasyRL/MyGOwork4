//MVC--Controller

package user_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/user"
	"bibi/pkg/errno"
	"bibi/pkg/utils"
)

func (s *UserService) Register(req *user.RegisterReq) (*db.User, error) {
	if len(req.Username) < 4 /*||len(req.Password)<8*/ {
		return nil, errno.ParamError
	}

	PwdDigest := utils.SetPassword(req.Password)
	userModel := &db.User{
		UserName: req.Username,
		Password: PwdDigest,
	}
	return db.Register(s.ctx, userModel)
}

func (s *UserService) Login(req *user.LoginReq) (*db.User, error) {
	userModel := &db.User{
		UserName: req.Username,
		Password: req.Password,
	}
	return db.Login(s.ctx, userModel)
}

func (s *UserService) Info(id int64) (*db.User, error) {
	userModel := &db.User{
		ID: id,
	}
	return db.QueryUserByID(userModel)
}
