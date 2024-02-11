package cache

import (
	rds "bibi/biz/mw/redis"
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func IsVideoLikeExist(ctx context.Context, videoId int64, uid int64) (bool, error) {
	r := NewRedisService(ctx, rds.RLike)
	return r.IsExist(i64ToStr(uid)+likeSuffix, i64ToStr(videoId))
}

func AddVideoLikeCount(ctx context.Context, videoId int64, uid int64) error {
	rLike := NewRedisService(ctx, rds.RLike)
	rVideo := NewRedisService(ctx, rds.RVideo)
	if err := rLike.Add(i64ToStr(uid)+likeSuffix, i64ToStr(videoId)); err != nil {
		return err
	}
	if err := rVideo.Increase(i64ToStr(videoId) + countSuffix); err != nil {
		return err
	}
	//设置count过期，在高压情况下可直接访问缓存
	if err := rVideo.Expire(i64ToStr(videoId) + countSuffix); err != nil {
		return err
	}
	return nil
}

func GetVideoLikeCount(ctx context.Context, videoId int64) (bool, int64, error) {
	rVideo := NewRedisService(ctx, rds.RVideo)
	v, err := rVideo.Get(i64ToStr(videoId) + countSuffix)
	if err == redis.Nil { //已过期
		return false, 0, nil
	}
	if err != nil {
		return true, 114514, err
	}
	cnt, _ := strconv.ParseInt(v, 10, 64)
	return true, cnt, nil
}

func DelVideoLikeCount(ctx context.Context, videoId int64, uid int64) error {
	rLike := NewRedisService(ctx, rds.RLike)
	rVideo := NewRedisService(ctx, rds.RVideo)
	if err := rLike.Del(i64ToStr(uid)+likeSuffix, i64ToStr(videoId)); err != nil {
		return err
	}
	if err := rVideo.Decrease(i64ToStr(videoId) + countSuffix); err != nil {
		return err
	}
	if err := rVideo.Expire(i64ToStr(videoId) + countSuffix); err != nil {
		return err
	}
	return nil
}
