package middleware

import (
	"base-system-backend/constants/code"
	"base-system-backend/constants/errmsg"
	"base-system-backend/model/common/response"
	"github.com/gin-gonic/gin"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	// NewBucket 默认的令牌桶，fillInterval 指每过多长时间向桶里放一个令牌，capacity 是桶的容量，超过桶容量的部分会被直接丢弃。桶初始是满的。
	// NewBucketWithRate 会按照提供的比例，每秒钟填充令牌数。例如 capacity 是100，而 rate 是 0.1，那么每秒会填充10个令牌。
	// NewBucketWithQuantum 和普通的 NewBucket() 的区别是，每次向桶中放令牌时，是放 quantum 个令牌，而不是一个令牌。
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			response.Error(c, code.RequestLimit, errmsg.RequestLimit, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}
