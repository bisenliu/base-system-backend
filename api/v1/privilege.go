package v1

import (
	"base-system-backend/enums/code"
	"base-system-backend/model/common/response"
	"base-system-backend/utils"
	"github.com/gin-gonic/gin"
)

type PrivilegeApi struct{}

func (PrivilegeApi) PrivilegeListApi(c *gin.Context) {
	results := map[string]interface{}{}
	//用户Id过滤
	userIdQueryString := c.Query("user_id")
	if userIdQueryString != "" {
		privilegeKeys, err, debugInfo := utils.PrivilegeUserIdFilter(userIdQueryString)
		if err != nil {
			response.Error(c, code.QueryFailed, err, debugInfo)
			return
		}
		results["results"] = privilegeKeys
		response.OK(c, results)
		return
	}
	//角色Id过滤
	roleIdQueryString := c.Query("role_id")
	if roleIdQueryString != "" {
		privilegeKeys, err, debugInfo := utils.PrivilegeRoleIdFilter(roleIdQueryString)
		if err != nil {
			response.Error(c, code.QueryFailed, err, debugInfo)
			return
		}
		results["results"] = privilegeKeys
		response.OK(c, results)
		return

	}
	privilegeKeys, err, debugInfo := privilegeService.PrivilegeListService()
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	results["results"] = privilegeKeys
	response.OK(c, results)
	return
}
