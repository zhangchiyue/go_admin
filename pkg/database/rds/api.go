package rds

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"time"
)

func defaultTimeoutCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), _ins.defaultTimeout)
}

func Get(key string) (*redis.StringCmd, error) {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()
	cmd := _ins.client.Get(ctx, key)
	if errors.Is(cmd.Err(), redis.Nil) { // 不存在
		return nil, nil
	}
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return cmd, nil
}
func Set(key string, data interface{}, expire ...time.Duration) error {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()

	dataExpire := _ins.defaultExpire
	if len(expire) > 0 {
		dataExpire = expire[0]
	}

	cmd := _ins.client.Set(ctx, key, data, dataExpire)
	return cmd.Err()
}

func SetNX(key string, data interface{}, expire ...time.Duration) (bool, error) {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()

	dataExpire := _ins.defaultExpire
	if len(expire) > 0 {
		dataExpire = expire[0]
	}

	cmd := _ins.client.SetNX(ctx, key, data, dataExpire)
	return cmd.Val(), cmd.Err()
}

func HGetAll(key string) (map[string]string, error) {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()

	cmd := _ins.client.HGetAll(ctx, key)
	return cmd.Val(), cmd.Err()
}

func SetString(key string, data string, expire ...time.Duration) error {
	return Set(key, data, expire...)
}

func GetString(key string) (string, error) {
	cmd, err := Get(key)
	if err != nil || cmd == nil {
		return "", err
	}
	return cmd.Val(), nil
}

func Del(key string) (bool, error) {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()
	cmd := _ins.client.Del(ctx, key)
	return cmd.Val() > 0, cmd.Err()
}

func Incr(key string) (int64, error) {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()
	res := _ins.client.Incr(ctx, key)
	return res.Val(), res.Err()
}
func HIncrBy(key, field string, incr int64) (int64, error) {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()

	cmd := _ins.client.HIncrBy(ctx, key, field, incr)
	return cmd.Val(), cmd.Err()
}

func Expire(key string, tm time.Duration) (bool, error) {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()
	res := _ins.client.Expire(ctx, key, tm)
	return res.Val(), res.Err()
}

func ExpireAt(key string, tm time.Time) (bool, error) {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()
	res := _ins.client.ExpireAt(ctx, key, tm)
	return res.Val(), res.Err()
}

func SAdd(key string, val ...interface{}) (uint64, error) {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()
	res := _ins.client.SAdd(ctx, key, val...)
	return res.Uint64()
}

func SMembers(key string) ([]string, error) {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()
	res := _ins.client.SMembers(ctx, key)
	return res.Val(), res.Err()
}

func SCard(key string) (int64, error) {
	ctx, cancel := defaultTimeoutCtx()
	defer cancel()
	res := _ins.client.SCard(ctx, key)
	return res.Val(), res.Err()
}
