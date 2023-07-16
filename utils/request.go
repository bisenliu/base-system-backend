package utils

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/constants/table"
	userEnum "base-system-backend/constants/user"
	"base-system-backend/global"
	"base-system-backend/model/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetLoginIp(c *gin.Context) (requestIp string) {
	if c.ClientIP() == "" {
		requestIp = c.RemoteIP()
	} else {
		requestIp = c.ClientIP()
	}
	// 来自CDN的回源，可能有多个IP
	if ok := strings.Contains(requestIp, ","); ok {
		requestIp = strings.Split(requestIp, ",")[0]
		return
	}
	return
}

func GetCurrentUser(c *gin.Context) (user *user.User, err error, debugInfo interface{}) {
	userId, ok := c.Get(userEnum.CtxUserIdKey)
	if !ok {
		return nil, fmt.Errorf("ctx用户%w", errmsg.NotFound), "ctx user not found"
	}

	if err = global.DB.Table(table.User).First(&user, "id = ?", userId).Error; err != nil {
		return nil, fmt.Errorf("用户%w", errmsg.QueryFailed), err.Error()
	}
	return
}
