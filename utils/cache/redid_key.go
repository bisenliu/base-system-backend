package cache

import "base-system-backend/global"

const KeyToken = "token:%s:" // token前级

func getRedisKey(key string) string {
	return global.CONFIG.Redis.Prefix + key
}
