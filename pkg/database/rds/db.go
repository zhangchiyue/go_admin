package rds

import (
	"adx-admin/pkg/configer"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	defaultExpire       = 30 * 24 * time.Hour
	defaultDialTimeout  = 30 * time.Second
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
	defaultMaxRetire    = 3
)

var _ins *RedisDB

type RedisDB struct {
	client         redis.UniversalClient
	defaultTimeout time.Duration
	defaultExpire  time.Duration
}

func NewRedisDB(config configer.Config) *RedisDB {
	client := redis.NewUniversalClient(getRedisConfigWithDefault(config.RedisServerConfig))
	_ins = &RedisDB{client: client, defaultTimeout: config.RedisServerConfig.BaseTimeout * time.Second, defaultExpire: defaultExpire}
	return _ins
}

func getRedisConfigWithDefault(config configer.RedisServerConfig) *redis.UniversalOptions {
	opt := &redis.UniversalOptions{}
	opt.Addrs = config.Addrs
	opt.Password = config.PassWord
	opt.DialTimeout = defaultDialTimeout
	opt.WriteTimeout = defaultWriteTimeout
	opt.ReadTimeout = defaultReadTimeout
	opt.MaxRetries = defaultMaxRetire
	if config.DialTimeout > 0 {
		opt.DialTimeout = config.DialTimeout
	}
	if config.WriteTimeout > 0 {
		opt.WriteTimeout = config.WriteTimeout
	}
	if config.ReadTimeout > 0 {
		opt.ReadTimeout = config.ReadTimeout
	}
	if config.MaxRetries > 0 {
		opt.MaxRedirects = config.MaxRetries
	}
	return opt
}
