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

// GetLoginIp
//  @Description: 获取当前登陆用户 IP
//  @param c 上下文信息
//  @return requestIp IP

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

// GetCurrentUser
//  @Description: 获取当前登陆用户实例
//  @param c 上下文信息
//  @return user 登陆用户
//  @return err 查询失败异常
//  @return debugInfo 错误调试信息

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
