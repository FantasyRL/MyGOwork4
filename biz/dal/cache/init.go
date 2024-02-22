package cache

import (
	"bibi/pkg/conf"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

const (
	likeSuffix          = ":like"
	followerCountSuffix = ":follower_counts"
	friendCountSuffix   = ":friend_counts"
	commentSuffix       = ":comment"
	followerSuffix      = ":follower"
	receiveSuffix       = ":receive"

	videoExpTime   = time.Hour * 1 //到期自动移除k-v
	likeExpTime    = time.Minute * 10
	commentExpTime = time.Minute * 10
	followExpTime  = time.Minute
	messageExpTime = time.Hour * 24 * 7 ////7天漫游(?

	videoLikeZset         = "video_likes"
	videoCommentCountZset = "video_comment_counts"
	videoCommentZset      = "video_comments"
	followerCountZset     = "follower_counts"
	followingCountZset    = "following_counts"
	friendCountZset       = "friend_counts"
)

var (
	rLike    *redis.Client
	rComment *redis.Client
	rFollow  *redis.Client
)

func Init() {
	rLike = redis.NewClient(&redis.Options{
		Addr:       conf.RedisAddr,
		ClientName: "Like",
		DB:         0,
	})
	rComment = redis.NewClient(&redis.Options{
		Addr:       conf.RedisAddr,
		ClientName: "Comment",
		DB:         1,
	})
	rFollow = redis.NewClient(&redis.Options{
		Addr:       conf.RedisAddr,
		ClientName: "Follow",
		DB:         2,
	})
}
func i64ToStr(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}
