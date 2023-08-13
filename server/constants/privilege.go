package constants

type PrivilegeKey string

// 自定义权限 key
const (
	SystemPrivilege    PrivilegeKey = "system_privilege"
	SystemManage       PrivilegeKey = "system_manage"
	LogManage          PrivilegeKey = "log_manage"
	OperateLogList     PrivilegeKey = "operate_log_list"
	OperateLogDownload PrivilegeKey = "operate_log_download"
	RoleManage         PrivilegeKey = "role_manage"
	RoleList           PrivilegeKey = "role_list"
	RoleDetail         PrivilegeKey = "role_detail"
	RoleCreate         PrivilegeKey = "role_create"
	RoleUpdate         PrivilegeKey = "role_update"
	RoleDelete         PrivilegeKey = "role_delete"
	RoleUnbind         PrivilegeKey = "role_unbind"
	PrivilegeManage    PrivilegeKey = "privilege_manage"
	PrivilegeList      PrivilegeKey = "privilege_list"
	PrivilegeSet       PrivilegeKey = "privilege_set"
	AccountManage      PrivilegeKey = "account_manage"
	AccountList        PrivilegeKey = "account_list"
	AccountCreate      PrivilegeKey = "account_create"
	ResetPwdOther      PrivilegeKey = "reset_pwd_other"
	ChangeStatusOther  PrivilegeKey = "change_status_other"
	AccountDetailOther PrivilegeKey = "account_detail_other"
	AccountUpdateOther PrivilegeKey = "account_update_other"
)

// Desc
//  @Description: 权限描述
//  @receiver p 接收者
//  @return string	描述

func (p PrivilegeKey) Desc() string {
	privilegeMap := map[PrivilegeKey]string{
		SystemPrivilege:    "系统权限",
		SystemManage:       "系统管理",
		LogManage:          "日志管理",
		OperateLogList:     "查询操作日志",
		OperateLogDownload: "下载操作日志",
		RoleManage:         "角色管理",
		RoleList:           "查询角色列表",
		RoleDetail:         "查询角色详情",
		RoleCreate:         "创建角色",
		RoleUpdate:         "修改角色",
		RoleDelete:         "删除角色",
		RoleUnbind:         "解绑角色",
		PrivilegeManage:    "权限管理",
		PrivilegeList:      "查询权限列表",
		PrivilegeSet:       "修改角色权限",
		AccountManage:      "账号管理",
		AccountList:        "查询账号列表",
		AccountCreate:      "添加账号",
		ResetPwdOther:      "重置指定账号密码",
		ChangeStatusOther:  "修改指定账号状态",
		AccountDetailOther: "查询指定账号信息",
		AccountUpdateOther: "修改指定账号信息",
	}
	return privilegeMap[p]
}

// String
//  @Description: 权限类型转字符串
//  @receiver p 接收者
//  @return string 字符串格式权限 key

func (p PrivilegeKey) String() string {
	return string(p)
}
