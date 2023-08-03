package cache

import (
	"base-system-backend/global"
	"context"
	"github.com/go-redis/redis/v8"
)

// newRdb
//
//	@Description: 选择对应库的 Redis
//	@param databaseName 库名 1-13
//	@return *redis.Client 连接对象

func newRdb(databaseName int) *redis.Client {
	rdb := global.REDIS
	rdb.Do(context.Background(), "SELECT", databaseName)
	return rdb
}
