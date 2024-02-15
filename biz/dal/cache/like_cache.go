package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

// IsUserLikeCacheExist 用户点赞是否存在于redis
func IsUserLikeCacheExist(ctx context.Context, uid int64) (bool, error) {
	ok, err := r.Exists(ctx, i64ToStr(uid)+likeSuffix).Result()
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

// IsVideoLikeExist 用户是否点赞了该视频
func IsVideoLikeExist(ctx context.Context, videoId int64, uid int64) (bool, error) {
	return r.SIsMember(ctx, i64ToStr(uid)+likeSuffix, i64ToStr(videoId)).Result()
}

// AddUserLikeVideo 仅添加用户点赞
func AddUserLikeVideo(ctx context.Context, videoId int64, uid int64) error {
	tx := r.TxPipeline()
	if err := tx.SAdd(ctx, i64ToStr(uid)+likeSuffix, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, i64ToStr(uid), likeExpTime).Err(); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

// AddVideoLikeCount 添加用户点赞、增加视频点赞量
func AddVideoLikeCount(ctx context.Context, videoId int64, uid int64) error {
	//管线很快，但组装命令过多会导致网络阻塞
	tx := r.TxPipeline()
	if err := tx.SAdd(ctx, i64ToStr(uid)+likeSuffix, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	if err := tx.ZIncrBy(ctx, videoLikeZset, 1, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	//刷新缓存时间
	if err := tx.Expire(ctx, videoLikeZset, videoExpTime).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, i64ToStr(uid), likeExpTime).Err(); err != nil {
		return err
	}

	if _, err := tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

// GetVideoLikeCount 获取视频点赞量
func GetVideoLikeCount(ctx context.Context, videoId int64) (bool, int64, error) {
	//获取元素的score
	v, err := r.ZScore(ctx, videoLikeZset, i64ToStr(videoId)).Result()
	if err == redis.Nil { //已过期
		return false, 0, nil
	}
	if err != nil {
		return true, 114514, err
	}
	cnt := int64(v)
	return true, cnt, nil
}

// DelVideoLikeCount 删除用户点赞、减少视频点赞量
func DelVideoLikeCount(ctx context.Context, videoId int64, uid int64) error {
	tx := r.TxPipeline()
	if err := tx.SRem(ctx, i64ToStr(uid)+likeSuffix, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	if err := tx.ZIncrBy(ctx, videoLikeZset, -1, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, videoLikeZset, videoExpTime).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, i64ToStr(uid), likeExpTime).Err(); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

// GetUserLikeVideos 获取用户点赞过的视频ID
func GetUserLikeVideos(ctx context.Context, uid int64) ([]int64, error) {
	//SMembers获取所有成员
	vals, err := r.SMembers(ctx, i64ToStr(uid)+likeSuffix).Result()
	if err != nil {
		return nil, err
	}

	var videoIdList []int64
	for _, id := range vals {
		vid, _ := strconv.ParseInt(id, 10, 64)
		videoIdList = append(videoIdList, vid)
	}
	return videoIdList, nil
}

// AddLikeVideoList 将用户的所有点赞写入到redis
func AddLikeVideoList(ctx context.Context, videoIdList []int64, uid int64) error {
	var err error
	for _, videoId := range videoIdList {
		err = r.SAdd(ctx, i64ToStr(uid)+likeSuffix, i64ToStr(videoId)).Err()
	}
	if err != nil {
		return err
	}
	err = r.Expire(ctx, i64ToStr(uid), likeExpTime).Err()
	return err
}

// SetVideoLikeCounts 将视频点赞量写入redis
func SetVideoLikeCounts(ctx context.Context, videoId int64, likeCount int64) error {
	err := r.ZAdd(ctx, videoLikeZset, redis.Z{
		Score:  float64(likeCount),
		Member: i64ToStr(videoId),
	}).Err()
	if err != nil {
		return err
	}
	err = r.Expire(ctx, videoLikeZset, videoExpTime).Err()
	return err
}

// ListHotVideo 通过Zset列出点赞最多的视频
func ListHotVideo(ctx context.Context) ([]int64, error) {
	//降序选择前4位点赞最高返回
	res, err := r.ZRevRange(ctx, videoLikeZset, 0, 4).Result()
	if err != nil {
		return nil, err
	}
	var videoIdList []int64
	for _, id := range res {
		vid, _ := strconv.ParseInt(id, 10, 64)
		videoIdList = append(videoIdList, vid)
	}
	return videoIdList, nil
}
