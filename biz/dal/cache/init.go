package cache

import (
	"bibi/pkg/conf"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

const (
	likeSuffix            = ":like"
	countSuffix           = ":count"
	commentSuffix         = ":comment"
	videoExpTime          = time.Hour * 1 //到期自动移除k-v
	likeExpTime           = time.Minute * 10
	commentExpTime        = time.Minute * 10
	videoLikeZset         = "video_likes"
	videoCommentCountZset = "video_comment_counts"
	videoCommentZset      = "video_comments"
)

var (
	r        *redis.Client
	rComment *redis.Client
)

func Init() {
	r = redis.NewClient(&redis.Options{
		Addr: conf.RedisAddr,
		DB:   0,
	})
	rComment = redis.NewClient(&redis.Options{
		Addr: conf.RedisAddr,
		DB:   1,
	})
	//rVideo = redis.NewClient(&redis.Options{
	//	Addr: conf.RedisAddr,
	//	DB:   2,
	//})
}
func i64ToStr(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}
