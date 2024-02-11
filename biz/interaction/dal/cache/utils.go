package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type RdsService struct {
	ctx context.Context
	rdb *redis.Client
}

const (
	likeSuffix  = ":like"
	countSuffix = ":count"
	expTime     = time.Hour * 1 //到期自动移除k-v
)

func NewRedisService(ctx context.Context, c *redis.Client) RdsService {
	return RdsService{
		ctx: ctx,
		rdb: c,
	}
}
func i64ToStr(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}

// Expire set Expire time
func (s *RdsService) Expire(k string) error {
	tx := s.rdb.TxPipeline() //管线很快，但组装命令过多会导致网络阻塞
	err := tx.Expire(s.ctx, k, expTime).Err()
	tx.Exec(s.ctx)
	return err
}

// Add k-v
func (s *RdsService) Add(k string, v string) error {
	tx := s.rdb.TxPipeline()
	err := tx.SAdd(s.ctx, k, v).Err()

	tx.Exec(s.ctx)
	return err
}

// Del k-v
func (s *RdsService) Del(k string, v string) error {
	tx := s.rdb.TxPipeline()
	err := tx.SRem(s.ctx, k, v).Err()
	tx.Exec(s.ctx)
	return err
}

// Get by key to value
func (s *RdsService) Get(k string) (string, error) {
	tx := s.rdb.TxPipeline()
	v, err := tx.Get(s.ctx, k).Result()
	if err != nil {
		return "nothingError", err
	}
	tx.Exec(s.ctx)
	return v, err
}

// Increase add the count
func (s *RdsService) Increase(k string) error {
	tx := s.rdb.TxPipeline()
	//Incr()、IncrBy()都是操作数字，对数字进行增加的操作，incr是执行原子加1操作，incrBy是增加指定的数
	err := tx.Incr(s.ctx, k).Err()
	tx.Exec(s.ctx)
	return err
}

// Decrease reduce the count
func (s *RdsService) Decrease(k string) error {
	tx := s.rdb.TxPipeline()
	err := tx.Decr(s.ctx, k).Err()
	tx.Exec(s.ctx)
	return err
}

// IsExist check k-v if IsExist
func (s *RdsService) IsExist(k string, v string) (bool, error) {
	e, err := s.rdb.SIsMember(s.ctx, k, v).Result()
	return e, err
}

// Sum get the size of k-v
func (s *RdsService) Sum(k string) int64 {
	total, _ := s.rdb.SCard(s.ctx, k).Result()
	s.rdb.Expire(s.ctx, k, expTime)
	return total
}
