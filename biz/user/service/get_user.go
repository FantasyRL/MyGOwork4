package service

import (
	"bibi/biz/user/dal/db"
	db2 "bibi/biz/video/dal/db"
)

func (s *UserService) GetUserByVideo(video db2.Video) (*db.User, error) {
	userModel := &db.User{
		ID: video.Uid,
	}
	return db.QueryUserByID(s.ctx, userModel)
}
func (s *UserService) GetUserById() {

}
