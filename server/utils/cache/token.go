package cache

import (
	"base-system-backend/global"
	"context"
	"fmt"
	"strconv"
	"time"
)

// getTokenExpireTime
// @Description: 获取token过期时间
// @return expire 过期时间
func getTokenExpireTime() (expire time.Duration) {
	return time.Duration(global.CONFIG.Token.ExpiredTime * float64(time.Hour))
}
func SetToken(userID int64, token string) {
	expireTime := getTokenExpireTime()
	key := fmt.Sprintf(getRedisKey(KeyToken), strconv.FormatInt(userID, 10))
	rdb := global.REDIS
	rdb.Do(context.Background(), "SELECT", global.CONFIG.Redis.TokenDb)
	rdb.Set(context.Background(), key, token, expireTime)
}

func GetToken(userID int64) string {
	key := fmt.Sprintf(getRedisKey(KeyToken), strconv.FormatInt(userID, 10))
	rdb := global.REDIS
	rdb.Do(context.Background(), "SELECT", global.CONFIG.Redis.TokenDb)
	token, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return ""
	}
	return token
}

// FlushToken 刷新登录token时间
func FlushToken(userID int64) {
	key := fmt.Sprintf(getRedisKey(KeyToken), strconv.FormatInt(userID, 10))
	expireTime := getTokenExpireTime()
	rdb := global.REDIS
	rdb.Do(context.Background(), "SELECT", global.CONFIG.Redis.TokenDb)
	rdb.Expire(context.Background(), key, expireTime)
}

// DeleteToken 删除token
func DeleteToken(userID int64) {
	key := fmt.Sprintf(getRedisKey(KeyToken), strconv.FormatInt(userID, 10))
	rdb := global.REDIS
	rdb.Do(context.Background(), "SELECT", global.CONFIG.Redis.TokenDb)
	rdb.Del(context.Background(), key)
}
