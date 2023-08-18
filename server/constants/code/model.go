package code

type Model []string

// ModelMapping  模型映射
var ModelMapping = map[string]Model{
	"user":      User,
	"role":      Role,
	"privilege": Privilege,
	"log":       Log,
	"version":   Version,
	"captcha":   Captcha,
}

// 模型对应的描述
var (
	User      Model = []string{"100", "user", "用户管理"}
	Role      Model = []string{"101", "role", "角色管理"}
	Privilege Model = []string{"102", "privilege", "权限管理"}
	Log       Model = []string{"103", "log", "日志管理"}
	Version   Model = []string{"104", "version", "版本管理"}
	Captcha   Model = []string{"105", "captcha", "滑块管理"}
	Unknown   Model = []string{"999", "unknown", "未知"}
)

// Choices
//  @Description: 根据 url 前缀获取对应的模型
//  @receiver m 接收者
//  @param prefix url 前缀
//  @return Model 模型

func (m Model) Choices(prefix string) Model {
	info, ok := ModelMapping[prefix]
	if !ok {
		info = Unknown
	}
	return info
}

// Code
//  @Description: 获取模型码
//  @receiver m 接收者
//  @return string 模型码

func (m Model) Code() string {
	return m[0]
}

// Key
//  @Description: 获取模型 key
//  @receiver m 接收者
//  @return string key

func (m Model) Key() string {
	return m[1]
}

// Desc
//  @Description: 获取模型描述
//  @receiver m 接收者
//  @return string 模型描述

func (m Model) Desc() string {
	return m[2]
}
