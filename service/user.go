package service

import (
	"base-system-backend/enums/errmsg"
	"base-system-backend/enums/table"
	"base-system-backend/global"
	"base-system-backend/model/user/request"
	"base-system-backend/model/user/response"
	"base-system-backend/utils"
	"base-system-backend/utils/orm"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserService struct{}

// UserListService
//
//	@Description: 用户列表
//	@param c 上下文信息
//	@param params 查询参数
//	@return userList 用户列表
//	@return err 查询失败异常
//	@return debugInfo 错误调试信息

func (receiver UserService) UserListService(c *gin.Context, params *request.UserFilter) (userList *response.UserList, err error, debugInfo interface{}) {
	// 过滤
	filter := make(map[string]map[string]string)
	// 账号/姓名
	if params.Name != "" {
		filter["LIKE"] = map[string]string{
			"name":    fmt.Sprintf("%%%s", params.Name),
			"account": fmt.Sprintf("%%%s", params.Name),
		}
	}
	// 状态
	if c.Query("status") != "" {
		filter["AND"] = map[string]string{
			"status": c.Query("status"),
		}
	}
	userList = new(response.UserList)
	if err = global.DB.Table(table.User).
		Scopes(orm.Paginate(params.Page, params.PageSize)).
		Scopes(orm.Where(filter)).
		Order("id").
		Find(&userList.Results).
		Limit(-1).Offset(-1).Count(&userList.TotalCount).Error; err != nil {
		return nil, fmt.Errorf("用户列表%w", errmsg.QueryFailed), err.Error()
	}
	for index, u := range userList.Results {
		privilegeKeys, userRoleIds, err, debugInfo := utils.GetPrivilegeKeysByUserId(u.Id)
		if err != nil {
			return nil, err, debugInfo
		}
		userList.Results[index].PrivilegeList = privilegeKeys
		userList.Results[index].RoleIds = userRoleIds
	}
	userList.GetPageInfo(&userList.PageInfo, params.Page, params.PageSize)
	return

}
