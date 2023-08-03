package cache

import "base-system-backend/global"

const KeyToken = "token:%s:" // token 前缀

// getRedisKey
//  @Description: 获取 redis 前缀
//  @param key 不同类型 key
//  @return string 前缀

func getRedisKey(key string) string {
	return global.CONFIG.Redis.Prefix + key
}
