package rds

import (
	"context"
	"fmt"
)

func (db *RedisDB) OnInit() error {
	ctx, cancel := context.WithTimeout(context.Background(), db.defaultTimeout)
	defer cancel()
	pong := db.client.Ping(ctx)
	if pong == nil || pong.Err() != nil {
		var err error
		if pong != nil {
			err = pong.Err()
		}
		return fmt.Errorf("ping redis server failed: %v", err)
	}
	return nil
}

func (db *RedisDB) OnDestroy() error {
	return db.stop()
}

func (db *RedisDB) Run(closeSig chan bool) {
	<-closeSig
}
func (db *RedisDB) stop() error {
	if err := db.client.Close(); err != nil {
		return err
	}
	return nil
}

func (db *RedisDB) Name() string {
	return "redisDB"
}
