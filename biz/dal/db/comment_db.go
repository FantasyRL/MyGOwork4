package db

import (
	"bibi/biz/model/interaction"
	"bibi/dao"
	"errors"
	"gorm.io/gorm"
	"time"
)

//go:generate msgp -tests=false
type Comment struct {
	ID        int64          `msg:"id"`
	VideoID   int64          `msg:"video_id"`
	Uid       int64          `msg:"uid"`
	Content   string         `msg:"content"`
	CreatedAt time.Time      `msg:"publish_time"`
	UpdatedAt time.Time      `msg:"-"`             //ignore
	DeletedAt gorm.DeletedAt `sql:"index" msg:"-"` //ignore
}

func IsCommentExist(commentModel *interaction.Comment) (bool, error) {
	var comment = &Comment{
		VideoID: commentModel.VideoID,
		Uid:     commentModel.User.ID,
		Content: commentModel.Content,
	}
	err := dao.DB.Model(Comment{}).Take(comment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func CreateComment(commentModel *interaction.Comment) (*Comment, error) {
	var comment = &Comment{
		VideoID: commentModel.VideoID,
		Uid:     commentModel.User.ID,
		Content: commentModel.Content,
	}
	if err := dao.DB.Model(Comment{}).Create(comment).Error; err != nil {
		return &Comment{}, err
	}
	return comment, nil
}

func DeleteComment(commentModel *interaction.Comment) (*Comment, error) {
	var comment = &Comment{
		VideoID: commentModel.VideoID,
		Uid:     commentModel.User.ID,
		Content: commentModel.Content,
	}
	if err := dao.DB.Model(Comment{}).Take(comment).Delete(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func GetCommentCount(videoId int64) (count int64, err error) {
	if err = dao.DB.Model(Comment{}).Where("video_id = ?", videoId).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}
