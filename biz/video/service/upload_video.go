package service

import (
	"bibi/biz/video/dal/db"
	"bibi/pkg/conf"
	"bytes"
	"log"
)

func (s *VideoService) UploadCover(cover []byte, name string) error {
	coverReader := bytes.NewReader(cover)
	err := s.bucket.PutObject(conf.OSSConf.MainDirectory+"/video/"+name, coverReader)
	if err != nil {
		log.Fatalf("upload file error:%video\n", err)
	}
	return err
}

func (s *VideoService) UploadVideo(video []byte, name string) error {
	videoReader := bytes.NewReader(video)
	err := s.bucket.PutObject(conf.OSSConf.MainDirectory+"/video/"+name, videoReader)
	if err != nil {
		log.Fatalf("upload file error:%video\n", err)
	}
	return err
}

func (s *VideoService) PutVideo(video *db.Video) (*db.Video, error) {
	return db.CreateVideo(s.ctx, video)
}
