package response

import (
	"base-system-backend/constants/login"
	"base-system-backend/constants/user"
	"base-system-backend/model/common/field"
	"base-system-backend/model/common/response"
)

type Token struct {
	Token string `json:"token"`
}

type LoginSuccess struct {
	UserDetail
	Token
}

//  UserDetail
//  @Description: 用户详情

type UserDetail struct {
	Id            int64             `json:"id"`
	Gender        int               `json:"gender"`
	IsSystem      bool              `json:"is_system"`
	Account       string            `json:"account"`
	Phone         *string           `json:"phone"`
	Email         *string           `json:"email"`
	Name          string            `json:"name"`
	IdCard        *string           `json:"id_card"`
	Avatar        *string           `json:"avatar"`
	Status        user.AccStatus    `json:"status"`
	LoginType     login.Type        `json:"login_type"`
	CreateTime    *field.CustomTime `json:"create_time"  swaggertype:"integer"`
	LastTime      *field.CustomTime `json:"last_time"  swaggertype:"integer"`
	PrivilegeList []string          `json:"privilege_list" gorm:"-"`
	RoleIds       []int64           `json:"role_ids" gorm:"-"`
}

//
//  UserList
//  @Description: 用户列表

type UserList struct {
	response.PageInfo
	Results []UserDetail `json:"results" form:"results"`
}
