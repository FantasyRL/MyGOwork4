//MVC--Controller

package service

import (
	"bibi/biz/model/user"
	db2 "bibi/biz/user/dal/db"
	"bibi/pkg/conf"
	"bibi/pkg/errno"
	"bibi/pkg/utils"
	"bytes"
	"log"
	"strconv"
)

func (s *UserService) Register(req *user.RegisterReq) (*db2.User, error) {
	if len(req.Username) < 4 /*||len(req.Password)<8*/ {
		return nil, errno.ParamError
	}

	PwdDigest := utils.SetPassword(req.Password)
	userModel := &db2.User{
		Name: req.Username,
		Pwd:  PwdDigest,
	}
	return db2.Register(s.ctx, userModel)
}

func (s *UserService) Login(req *user.LoginReq) (*db2.User, error) {
	userModel := &db2.User{
		Name: req.Username,
		Pwd:  req.Password,
	}
	return db2.Login(s.ctx, userModel)
}

func (s *UserService) Info(req *user.InfoReq) (*db2.User, error) {
	userModel := &db2.User{
		ID: req.UserID,
	}
	return db2.QueryUserByID(s.ctx, userModel)
}

func (s *AvatarService) UploadAvatar(req *user.AvatarReq, id int64) error {
	avatarReader := bytes.NewReader(req.AvatarFile)
	err := s.bucket.PutObject(conf.OSSConf.MainDirectory+"/"+strconv.FormatInt(id, 10)+".jpg", avatarReader)
	if err != nil {
		log.Fatalf("upload file error:%v\n", err)
	}
	return err
}

func (s *AvatarService) PutAvatar(id int64, avatarUrl string) (*db2.User, error) {
	userModel := &db2.User{
		ID:     id,
		Avatar: avatarUrl,
	}
	return db2.PutAvatar(s.ctx, userModel)
}
