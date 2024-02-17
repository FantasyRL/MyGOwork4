package cache

import (
	"bibi/biz/dal/db"
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func AddVideoComment(ctx context.Context, comment *db.Comment) error {
	tx := rComment.TxPipeline()
	//msgp序列化结构体
	MalComment, err := (*comment).MarshalMsg(nil)
	if err != nil {
		return err
	}
	//Zset 相比hash可以用score来排序
	if err := tx.ZAdd(ctx, i64ToStr(comment.VideoID)+commentSuffix, redis.Z{
		Score:  float64(comment.ID),
		Member: MalComment,
	}).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, i64ToStr(comment.VideoID)+commentSuffix, commentExpTime).Err(); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func SetVideoComments(ctx context.Context, comments []db.Comment, videoId int64) error {
	var err error
	MalComments := make([]redis.Z, len(comments))
	for i, comment := range comments {
		var MalComment []byte
		MalComment, err = (comment).MarshalMsg(nil)

		MalComments[i] = redis.Z{
			Score:  float64(comment.ID),
			Member: MalComment,
		}
	}
	if err != nil {
		return err
	}
	tx := rComment.TxPipeline()
	//MalComments...可存储整个切片
	if err = tx.ZAdd(ctx, i64ToStr(videoId)+commentSuffix, MalComments...).Err(); err != nil {
		return err
	}
	if err = tx.Expire(ctx, i64ToStr(videoId)+commentSuffix, commentExpTime).Err(); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func DelVideoComment(ctx context.Context, comment *db.Comment) error {
	score := strconv.FormatInt(comment.ID, 10)
	return rComment.ZRemRangeByScore(ctx, i64ToStr(comment.VideoID)+commentSuffix, score, score).Err()

}

func GetVideoComments(ctx context.Context, videoId int64) (comments []db.Comment, err error) {
	//按时间排序获取所有评论
	malComments, err := rComment.ZRevRange(ctx, i64ToStr(videoId)+commentSuffix, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	for _, malComment := range malComments {
		var comment db.Comment
		_, err = comment.UnmarshalMsg([]byte(malComment))
		comments = append(comments, comment)
	}
	if err != nil {
		return nil, err
	}
	return
}

func GetVideoCommentCount(ctx context.Context, videoId int64) (bool, int64, error) {
	v, err := r.ZScore(ctx, videoCommentCountZset, i64ToStr(videoId)).Result()
	if err == redis.Nil { //已过期
		return false, 0, nil
	}
	if err != nil {
		return true, 114514, err
	}
	cnt := int64(v)
	return true, cnt, nil
}

func IncrVideoCommentCount(ctx context.Context, videoId int64) error {
	tx := r.TxPipeline()
	if err := tx.ZIncrBy(ctx, videoCommentCountZset, 1, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, videoCommentCountZset, videoExpTime).Err(); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func DecrVideoCommentCount(ctx context.Context, videoId int64) error {
	tx := r.TxPipeline()
	if err := tx.ZIncrBy(ctx, videoCommentCountZset, -1, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, videoCommentCountZset, videoExpTime).Err(); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func SetVideoCommentCount(ctx context.Context, videoId int64, count int64) error {
	tx := r.TxPipeline()
	if err := tx.ZAdd(ctx, videoCommentCountZset, redis.Z{
		Score:  float64(count),
		Member: i64ToStr(videoId),
	}).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, videoCommentCountZset, videoExpTime).Err(); err != nil {
		return err
	}
	_, err := tx.Exec(ctx)
	return err
}
