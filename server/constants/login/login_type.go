package login

type Type int

// 登陆类型
const (
	AccPwdLogin   Type = iota // 账号密码登陆
	PhoneLogin                // 手机号登陆
	KeycloakLogin             // keycloak登陆

)

// IsValid
//  @Description: 登陆类型枚举校验,配合自定义 validator
//  @receiver t 接收者
//  @return bool 是否校验通过

func (t Type) IsValid() bool {
	switch t {
	case AccPwdLogin, PhoneLogin, KeycloakLogin:
		return true
	}
	return false
}

// MaxLoginFailedNum 最大登陆失败次数
const MaxLoginFailedNum = 5
