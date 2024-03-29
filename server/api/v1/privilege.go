package v1

import (
	"base-system-backend/constants/code"
	"base-system-backend/model/common/response"
	"base-system-backend/model/privilege/request"
	"base-system-backend/utils/validate"
	"github.com/gin-gonic/gin"
)

type PrivilegeApi struct{}

// PrivilegeListApi
// @Summary 权限列表
// @Description 权限列表
// @Tags PrivilegeApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param user_id query string false "用户 ID"
// @Param role_id query string false "角色 ID"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=response.PrivilegeList}
// @Router /privilege/list/ [get]
func (PrivilegeApi) PrivilegeListApi(c *gin.Context) {
	results := map[string]interface{}{}
	//用户Id过滤
	userId := c.Query("user_id")
	//角色Id过滤
	roleId := c.Query("role_id")
	privilegeKeys, err, debugInfo := privilegeService.PrivilegeListService(userId, roleId)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	results["results"] = privilegeKeys
	response.OK(c, results)
}

// RolePrivilegeUpdateApi
// @Summary 更新角色权限
// @Description 更新角色权限
// @Tags PrivilegeApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object body request.RolePrivilegeUpdate true "更新参数"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data
// @Router /privilege/:role_id/ [put]
func (PrivilegeApi) RolePrivilegeUpdateApi(c *gin.Context) {
	params := new(request.RolePrivilegeUpdate)
	if !validate.RequestParamsVerify(c, params) {
		return
	}
	roleId := c.Param("role_id")
	if err, debugInfo := privilegeService.RolePrivilegeUpdateService(roleId, params); err != nil {
		response.Error(c, code.UpdateFailed, err, debugInfo)
		return
	}
	response.OK(c, nil)
}
