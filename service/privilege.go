package service

import (
	"base-system-backend/enums/errmsg"
	"base-system-backend/enums/table"
	"base-system-backend/global"
	"base-system-backend/model/privilege/response"
	"base-system-backend/utils"
	"fmt"
)

type PrivilegeService struct{}

func (receiver PrivilegeService) PrivilegeListService() (res []*response.PrivilegeList, err error, debugInfo interface{}) {
	var privileges []*response.PrivilegeList
	if err = global.DB.Table(table.Privilege).Find(&privileges).Error; err != nil {
		return nil, fmt.Errorf("权限信息%w", errmsg.QueryFailed), err.Error()
	}
	res = utils.RecursionGetChildPrivilege(privileges, 1)
	return
}
