package cache

import (
	"bibi/biz/dal/db"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

func SetMessage(ctx context.Context, message *db.Message) error {
	tx := rMessage.TxPipeline()
	marshalMsg, err := (*message).MarshalMsg(nil)
	if err != nil {
		return err
	}
	if err = tx.ZAdd(ctx, i64ToStr(message.TargetId)+receiveSuffix, redis.Z{
		Score:  float64(message.ID),
		Member: marshalMsg,
	}).Err(); err != nil {
		return err
	}
	if err = tx.Expire(ctx, i64ToStr(message.TargetId)+receiveSuffix, messageExpTime).Err(); err != nil {
		return err
	}
	if _, err = tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func GetMessages(ctx context.Context, uid int64) (messages []db.Message, err error) {
	marshalMsgs, err := rMessage.ZRevRange(ctx, i64ToStr(uid)+receiveSuffix, 0, -1).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	for _, marshalMsg := range marshalMsgs {
		var message db.Message
		_, err = message.UnmarshalMsg([]byte(marshalMsg))
		messages = append(messages, message)
	}
	if err != nil {
		return nil, err
	}
	return
}
