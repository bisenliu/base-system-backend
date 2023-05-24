package v1

import (
	"base-system-backend/enums/code"
	Rsp "base-system-backend/model/common/response"
	"base-system-backend/model/user/request"
	"base-system-backend/utils/validate"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// UserListApi
// @Summary 用户列表
// @Description 用户列表
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query request.UserFilter false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=response.UserList}
// @Router /user/list/ [get]
func (UserApi) UserListApi(c *gin.Context) {
	params := new(request.UserFilter)
	if ok := validate.QueryParamsVerify(c, &params); !ok {
		return
	}
	userList, err, debugInfo := userService.UserListService(c, params)
	if err != nil {
		Rsp.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	Rsp.OK(c, userList)
	return
}
