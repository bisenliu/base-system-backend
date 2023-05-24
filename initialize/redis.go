package initialize

import (
	"base-system-backend/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
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

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic(fmt.Errorf("redis connect failed: %s", err))
	}
}

func CloseRedis() {
	_ = global.REDIS.Close()
}
