package v1

import (
	"base-system-backend/enums/code"
	"base-system-backend/model/common/response"
	"base-system-backend/model/role/request"
	"base-system-backend/utils/validate"
	"github.com/gin-gonic/gin"
)

type RoleApi struct{}

// RoleListApi
// @Summary 角色列表
// @Description 角色列表
// @Tags RoleApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object query request.RoleFilter false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=response.RoleList}
// @Router /role/list/ [get]
func (RoleApi) RoleListApi(c *gin.Context) {
	roleList, err, debugInfo := roleService.RoleListService(c)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
	}
	response.OK(c, map[string]interface{}{
		"results": roleList,
	})
}

// RoleCreateApi
// @Summary 角色添加
// @Description 角色添加
// @Tags RoleApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object body request.RoleCreate true "角色信息"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=response.Create}
// @Router /role/ [post]
func (RoleApi) RoleCreateApi(c *gin.Context) {
	params := new(request.RoleCreate)
	if ok := validate.RequestParamsVerify(c, params); !ok {
		return
	}
	if err, debugInfo := roleService.RoleCreateService(params); err != nil {
		response.Error(c, code.SaveFailed, err, debugInfo)
	}
	response.OK(c, map[string]interface{}{"id": params.Id})
}

// RoleDetailApi
// @Summary 角色详情
// @Description 角色详情
// @Tags RoleApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=role.Role}
// @Router /role/:role_id// [get]
func (RoleApi) RoleDetailApi(c *gin.Context) {
	roleId := c.Param("role_id")
	roleDetail, err, debugInfo := roleService.RoleDetailService(roleId)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
	}
	response.OK(c, *roleDetail)
}

// RoleUpdateApi
// @Summary 角色修改
// @Description 角色修改
// @Tags RoleApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object body request.RoleUpdate true "角色信息"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data
// @Router /role/:role_id/ [put]
func (RoleApi) RoleUpdateApi(c *gin.Context) {
	params := new(request.RoleUpdate)
	if ok := validate.RequestParamsVerify(c, params); !ok {
		return
	}
	roleId := c.Param("role_id")
	if err, debugInfo := roleService.RoleUpdateService(roleId, params); err != nil {
		response.Error(c, code.UpdateFailed, err, debugInfo)
	}
	response.OK(c, nil)
}

// RoleDeleteApi
// @Summary 角色删除
// @Description 角色删除
// @Tags RoleApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data
// @Router /role/:role_id/ [delete]
func (RoleApi) RoleDeleteApi(c *gin.Context) {
	roleId := c.Param("role_id")
	if err, debugInfo := roleService.RoleDeleteService(roleId); err != nil {
		response.Error(c, code.DeleteFailed, err, debugInfo)
	}
	response.OK(c, nil)
}
