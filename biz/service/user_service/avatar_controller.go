package user_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/user"
	"bibi/pkg/conf"
	"bytes"
	"log"
	"strconv"
)

func (s *AvatarService) UploadAvatar(req *user.AvatarReq, id int64) error {
	avatarReader := bytes.NewReader(req.AvatarFile)
	err := s.bucket.PutObject(conf.OSSConf.MainDirectory+"/"+strconv.FormatInt(id, 10)+".jpg", avatarReader)
	if err != nil {
		log.Fatalf("upload file error:%video\n", err)
	}
	return err
}

func (s *AvatarService) PutAvatar(id int64, avatarUrl string) (*db.User, error) {
	userModel := &db.User{
		ID:     id,
		Avatar: avatarUrl,
	}
	return db.PutAvatar(s.ctx, userModel)
}
