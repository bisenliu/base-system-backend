package initialize

import (
	"base-system-backend/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisCfg.Host, redisCfg.Port),
		Password: redisCfg.Password,  // no password set
		DB:       redisCfg.DefaultDb, // use default DB
		//PoolSize:     cfg.PoolSize,
		//MinIdleConns: cfg.MinIdleConns,
	})

	if pong, err := client.Ping(context.Background()).Result(); err != nil {
		panic(fmt.Errorf("redis connect failed: %w", err))
	} else {
		global.LOG.Info("redis connect info: ", zap.String("pong", pong))
		global.REDIS = client
	}
}

func CloseRedis() {
	_ = global.REDIS.Close()
}
