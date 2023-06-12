package v1

import (
	"base-system-backend/enums/code"
	"base-system-backend/model/common/response"
	"base-system-backend/model/privilege/request"
	"base-system-backend/utils"
	"base-system-backend/utils/validate"
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
		}
		results["results"] = privilegeKeys
		response.OK(c, results)
	}
	//角色Id过滤
	roleIdQueryString := c.Query("role_id")
	if roleIdQueryString != "" {
		privilegeKeys, err, debugInfo := utils.PrivilegeRoleIdFilter(roleIdQueryString)
		if err != nil {
			response.Error(c, code.QueryFailed, err, debugInfo)
		}
		results["results"] = privilegeKeys
		response.OK(c, results)

	}
	privilegeKeys, err, debugInfo := privilegeService.PrivilegeListService()
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
	}
	results["results"] = privilegeKeys
	response.OK(c, results)
}

func (PrivilegeApi) RolePrivilegeUpdateApi(c *gin.Context) {
	params := new(request.RolePrivilegeUpdate)
	if ok := validate.RequestParamsVerify(c, params); !ok {
		return
	}
	roleId := c.Param("role_id")
	if err, debugInfo := privilegeService.RolePrivilegeUpdateService(roleId, params); err != nil {
		response.Error(c, code.UpdateFailed, err, debugInfo)
	}
	response.OK(c, nil)
}
