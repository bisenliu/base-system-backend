package cache

import (
	"base-system-backend/global"
	"context"
	"github.com/go-redis/redis/v8"
)

func newRdb(databaseName int) *redis.Client {
	rdb := global.REDIS
	rdb.Do(context.Background(), "SELECT", databaseName)
	return rdb
}
