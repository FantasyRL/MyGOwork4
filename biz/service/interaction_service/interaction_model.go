package interaction_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/interaction"
	"bibi/biz/model/user"
	"context"
)

type InteractionService struct {
	ctx context.Context
}

func NewInteractionService(ctx context.Context) *InteractionService {
	return &InteractionService{ctx: ctx}
}

func BuildCommentResp(comment *db.Comment) *interaction.Comment {
	commenter, _ := db.QueryUserByID(&db.User{ID: comment.Uid})
	return &interaction.Comment{
		ID:          comment.ID,
		VideoID:     comment.VideoID,
		User:        BuildUserResp(commenter),
		Content:     comment.Content,
		PublishTime: comment.CreatedAt.Format("2006-01-02 15:01:04"),
	}
}

func BuildCommentsResp(comments []db.Comment) (commentsResp []*interaction.Comment) {
	for _, comment := range comments {
		commentsResp = append(commentsResp, BuildCommentResp(&comment))
	}
	return
}

func BuildUserResp(commenter *db.User) *user.User {
	return &user.User{
		ID:     commenter.ID,
		Name:   commenter.UserName,
		Avatar: commenter.Avatar,
	}
}
