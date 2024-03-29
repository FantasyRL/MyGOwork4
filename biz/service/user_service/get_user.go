package user_service

import (
	"bibi/biz/dal/db"
)

func (s *UserService) GetUserByVideo(video db.Video) (*db.User, error) {
	userModel := &db.User{
		ID: video.Uid,
	}
	return db.QueryUserByID(userModel)
}
func (s *UserService) GetUserByIdList(uidList []int64) ([]db.User, error) {
	return db.QueryUserByIDList(uidList)
}
