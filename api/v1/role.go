package v1

import (
	"base-system-backend/enums/code"
	"base-system-backend/model/common/response"
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
