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
	}
	response.OK(c, map[string]interface{}{
		"results": roleList,
	})
}

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

func (RoleApi) RoleDetailApi(c *gin.Context) {
	roleId := c.Param("role_id")
	roleDetail, err, debugInfo := roleService.RoleDetailService(roleId)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
	}
	response.OK(c, *roleDetail)
}

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

func (RoleApi) RoleDeleteApi(c *gin.Context) {
	roleId := c.Param("role_id")
	if err, debugInfo := roleService.RoleDeleteService(roleId); err != nil {
		response.Error(c, code.DeleteFailed, err, debugInfo)
	}
	response.OK(c, nil)
}
