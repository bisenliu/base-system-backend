package request

import (
	"base-system-backend/enums/user"
	"base-system-backend/model/common/request"
)

type UserFilter struct {
	Name   string         `json:"name" form:"name" binding:"" label:"用户名"`
	Status user.AccStatus `json:"status" form:"status" binding:"enum" label:"用户状态"`
	request.PageInfo
}
