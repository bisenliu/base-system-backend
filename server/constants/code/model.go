package code

type Model []string

var ModelMapping = map[string]Model{
	"user":      User,
	"role":      Role,
	"privilege": Privilege,
	"log":       Log,
	"version":   Version,
}

var (
	User      Model = []string{"100", "user", "用户管理"}
	Role      Model = []string{"101", "role", "角色管理"}
	Privilege Model = []string{"102", "privilege", "权限管理"}
	Log       Model = []string{"103", "log", "日志管理"}
	Version   Model = []string{"104", "version", "版本管理"}
	Unknown   Model = []string{"999", "unknown", "未知"}
)

func (m Model) Choices(prefix string) Model {
	info, ok := ModelMapping[prefix]
	if !ok {
		info = Unknown
	}
	return info
}

func (m Model) Code() string {
	return m[0]
}

func (m Model) Key() string {
	return m[1]
}

func (m Model) Desc() string {
	return m[2]
}
