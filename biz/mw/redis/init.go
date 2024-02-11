package redis

import (
	"bibi/pkg/conf"
	"github.com/redis/go-redis/v9"
)

var (
	RLike  *redis.Client
	RVideo *redis.Client
)

func Init() {
	RLike = redis.NewClient(&redis.Options{
		Addr: conf.RedisAddr,
		DB:   0,
	})
	//RSub = redis.NewClient(&redis.Options{
	//	Addr: conf.RedisAddr,
	//	DB:   1,
	//})
	RVideo = redis.NewClient(&redis.Options{
		Addr: conf.RedisAddr,
		DB:   2,
	})
}
