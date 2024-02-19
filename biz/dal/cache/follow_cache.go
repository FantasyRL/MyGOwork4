package cache

import (
	"bibi/biz/dal/db"
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func IsFollowedCacheExist(ctx context.Context, followedId int64) (bool, error) {
	ok, err := rFollow.Exists(ctx, i64ToStr(followedId)+followerSuffix).Result()
	if err != nil {
		//错误处理返回啥都一样
		return false, err
	}
	if ok > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func IsUserFollowExist(ctx context.Context, followerId int64, followedId int64) (bool, error) {
	return rFollow.SIsMember(ctx, i64ToStr(followedId)+followerSuffix, i64ToStr(followerId)).Result()
}

func SetFollowerList(ctx context.Context, uid int64, followList []db.Follow) (err error) {
	tx := rFollow.TxPipeline()
	for _, follow := range followList {
		err = tx.SAdd(ctx, i64ToStr(uid)+followerSuffix, i64ToStr(follow.Uid)).Err()
	}
	if err != nil {
		return
	}
	if err = tx.Expire(ctx, i64ToStr(uid)+followerSuffix, followExpTime).Err(); err != nil {
		return
	}
	_, err = tx.Exec(ctx)
	return
}

func SetFollowingList(ctx context.Context, uid int64, followList []db.Follow) (err error) {
	tx := rFollow.TxPipeline()
	for _, follow := range followList {
		err = tx.SAdd(ctx, i64ToStr(follow.FollowedId)+followerSuffix, i64ToStr(uid)).Err()
	}
	if err != nil {
		return
	}
	for _, follow := range followList {
		err = tx.Expire(ctx, i64ToStr(follow.FollowedId)+followerSuffix, followExpTime).Err()
	}

	_, err = tx.Exec(ctx)
	return
}

func SetFriendList(ctx context.Context, uid int64, followList []db.Follow) (err error) {
	tx := rFollow.TxPipeline()
	for _, follow := range followList {
		err = tx.SAdd(ctx, i64ToStr(uid)+followerSuffix, i64ToStr(follow.Uid)).Err()
		err = tx.SAdd(ctx, i64ToStr(follow.FollowedId)+followerSuffix, i64ToStr(uid)).Err()
	}
	if err != nil {
		return
	}
	if err = tx.Expire(ctx, i64ToStr(uid)+followerSuffix, followExpTime).Err(); err != nil {
		return
	}
	for _, follow := range followList {
		err = tx.Expire(ctx, i64ToStr(follow.FollowedId)+followerSuffix, followExpTime).Err()
	}
	_, err = tx.Exec(ctx)
	return
}

func SetFollowerCount(ctx context.Context, uid int64, count int64) (err error) {
	err = rFollow.ZAdd(ctx, followerCountZset, redis.Z{
		Score:  float64(count),
		Member: i64ToStr(uid),
	}).Err()
	if err != nil {
		return
	}
	err = rFollow.Expire(ctx, followerCountZset, followExpTime).Err()
	return
}

func SetFollowingCount(ctx context.Context, uid int64, count int64) (err error) {
	err = rFollow.ZAdd(ctx, followingCountZset, redis.Z{
		Score:  float64(count),
		Member: i64ToStr(uid),
	}).Err()
	if err != nil {
		return
	}
	err = rFollow.Expire(ctx, followingCountZset, followExpTime).Err()
	return
}

func SetFriendCount(ctx context.Context, uid int64, count int64) (err error) {
	err = rFollow.ZAdd(ctx, FriendCountZset, redis.Z{
		Score:  float64(count),
		Member: i64ToStr(uid),
	}).Err()
	if err != nil {
		return err
	}
	err = rFollow.Expire(ctx, FriendCountZset, followExpTime).Err()
	return
}

func AddFollower(ctx context.Context, followerId int64, followedId int64) (err error) {
	tx := rFollow.TxPipeline()
	if err = tx.SAdd(ctx, i64ToStr(followedId)+followerSuffix, i64ToStr(followerId)).Err(); err != nil {
		return err
	}
	if err = tx.ZIncrBy(ctx, followerCountZset, 1, i64ToStr(followedId)).Err(); err != nil {
		return err
	}
	if err = tx.Expire(ctx, i64ToStr(followedId)+followerSuffix, followExpTime).Err(); err != nil {
		return
	}
	if err = tx.Expire(ctx, followerCountZset, followExpTime).Err(); err != nil {
		return
	}
	_, err = tx.Exec(ctx)
	return
}

func DelFollower(ctx context.Context, followerId int64, followedId int64) (err error) {
	tx := rFollow.TxPipeline()
	if err = tx.SRem(ctx, i64ToStr(followedId)+followerSuffix, i64ToStr(followerId)).Err(); err != nil {
		return err
	}
	if err = tx.ZIncrBy(ctx, followerCountZset, -1, i64ToStr(followedId)).Err(); err != nil {
		return err
	}
	if err = tx.Expire(ctx, i64ToStr(followedId)+followerSuffix, followExpTime).Err(); err != nil {
		return
	}
	if err = tx.Expire(ctx, followerCountZset, followExpTime).Err(); err != nil {
		return
	}
	_, err = tx.Exec(ctx)
	return
}

func AddFollowing(ctx context.Context, followerId int64) (err error) {
	if err = rFollow.ZIncrBy(ctx, followingCountZset, 1, i64ToStr(followerId)).Err(); err != nil {
		return
	}
	if err = rFollow.Expire(ctx, followingCountZset, followExpTime).Err(); err != nil {
		return
	}
	return nil
}

func DelFollowing(ctx context.Context, followerId int64) (err error) {
	if err = rFollow.ZIncrBy(ctx, followingCountZset, 1, i64ToStr(followerId)).Err(); err != nil {
		return
	}
	if err = rFollow.Expire(ctx, followingCountZset, followExpTime).Err(); err != nil {
		return
	}
	return nil
}

func GetFollower(ctx context.Context, uid int64) ([]db.Follow, error) {
	vals, err := rFollow.SMembers(ctx, i64ToStr(uid)+followerSuffix).Result()
	if err != nil {
		return nil, err
	}
	var followerList []db.Follow
	for _, follower := range vals {
		followerId, _ := strconv.ParseInt(follower, 10, 64)
		followerList = append(followerList, db.Follow{
			Uid:        followerId,
			FollowedId: uid,
		})
	}
	return followerList, nil
}

func GetFollowerCount(ctx context.Context, uid int64) (bool, int64, error) {
	v, err := rFollow.ZScore(ctx, followerCountZset, i64ToStr(uid)).Result()
	if err == redis.Nil { //已过期
		return false, 0, nil
	}
	if err != nil {
		return true, 114514, err
	}
	cnt := int64(v)
	return true, cnt, nil
}

func GetFollowingCount(ctx context.Context, uid int64) (bool, int64, error) {
	v, err := rFollow.ZScore(ctx, followingCountZset, i64ToStr(uid)).Result()
	if err == redis.Nil { //已过期
		return false, 0, nil
	}
	if err != nil {
		return true, 114514, err
	}
	cnt := int64(v)
	return true, cnt, nil
}
