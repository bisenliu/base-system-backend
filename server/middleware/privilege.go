package middleware

import (
	"base-system-backend/constants"
	"base-system-backend/constants/code"
	"base-system-backend/constants/errmsg"
	"base-system-backend/model/common/response"
	"base-system-backend/utils"
	"base-system-backend/utils/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

// PrivilegeVerify
// @Description：权限校验中间件
// @param key 每个url对应的权限key
// @return gin.HandlerFunc

func PrivilegeVerify(key constants.PrivilegeKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err, debugInfo := utils.GetCurrentUser(c)
		if err != nil {
			response.Error(c, code.InvalidLogin, err, debugInfo)
			c.Abort()
			return
		}
		//获取角色ID列表
		privilegeKeys, _, err, debugInfo := utils.GetPrivilegeKeysByUserId(user.Id)
		if err != nil {
			response.Error(c, code.InvalidLogin, err, debugInfo)
			c.Abort()
			return
		}
		if !common.In(string(key), privilegeKeys) {
			response.Error(c, code.NotPermissions, fmt.Sprintf(errmsg.NotPrivilege.Error(), key.Desc()), debugInfo)
			c.Abort()
			return
		}
	}
}
