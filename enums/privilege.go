package enums

type PrivilegeKey string

const (
	UserListPrivilege PrivilegeKey = "user_list"
)

func (PrivilegeKey) PrivilegeKey(key PrivilegeKey) string {
	privilegeMap := map[PrivilegeKey]string{
		UserListPrivilege: "用户列表",
	}
	return privilegeMap[key]
}
