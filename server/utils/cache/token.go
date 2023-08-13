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

// SetToken
//  @Description: 缓存 token
//  @param userID 用户 ID
//  @param token token

func SetToken(userID int64, token string) {
	expireTime := getTokenExpireTime()
	key := fmt.Sprintf(getRedisKey(KeyToken), strconv.FormatInt(userID, 10))
	newRdb(global.CONFIG.Redis.TokenDb).Set(context.Background(), key, token, expireTime)
}

// GetToken
//  @Description: 获取缓存 token
//  @param userID 用户 ID
//  @return string token

func GetToken(userID int64) string {
	key := fmt.Sprintf(getRedisKey(KeyToken), strconv.FormatInt(userID, 10))
	token, err := newRdb(global.CONFIG.Redis.TokenDb).Get(context.Background(), key).Result()
	if err != nil {
		return ""
	}
	return token
}

// FlushToken
//  @Description: 刷新登陆token时间
//  @param userID

func FlushToken(userID int64) {
	key := fmt.Sprintf(getRedisKey(KeyToken), strconv.FormatInt(userID, 10))
	expireTime := getTokenExpireTime()
	newRdb(global.CONFIG.Redis.TokenDb).Expire(context.Background(), key, expireTime)
}

// DeleteToken
//  @Description: 删除token
//  @param userID userID

func DeleteToken(userID int64) {
	key := fmt.Sprintf(getRedisKey(KeyToken), strconv.FormatInt(userID, 10))
	newRdb(global.CONFIG.Redis.TokenDb).Del(context.Background(), key)
}
