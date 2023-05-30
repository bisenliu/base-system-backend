package v1

import (
	"base-system-backend/enums/code"
	"base-system-backend/model/common/response"
	"base-system-backend/model/role/request"
	"base-system-backend/utils/validate"
	"github.com/gin-gonic/gin"
)

type RoleApi struct{}

func (RoleApi) RoleListApi(c *gin.Context) {
	roleList, err, debugInfo := roleService.RoleListService(c)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	response.OK(c, map[string]interface{}{
		"results": roleList,
	})
	return
}

func (RoleApi) RoleCreateApi(c *gin.Context) {
	params := new(request.RoleCreate)
	if ok := validate.RequestParamsVerify(c, params); !ok {
		return
	}
	if err, debugInfo := roleService.RoleCreateService(params); err != nil {
		response.Error(c, code.SaveFailed, err, debugInfo)
		return
	}
	response.OK(c, map[string]interface{}{"id": params.Id})
	return
}

func (RoleApi) RoleDetailApi(c *gin.Context) {
	roleId := c.Param("role_id")
	roleDetail, err, debugInfo := roleService.RoleDetailService(roleId)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	response.OK(c, *roleDetail)
	return
}
