package login

type LoginType int

const (
	AccPwdLogin   = iota // 账号密码登录
	PhoneLogin           // 手机号登录
	KeycloakLogin        // keycloak登录

)

func (t LoginType) IsValid() bool {
	switch t {
	case AccPwdLogin, PhoneLogin, KeycloakLogin:
		return true
	}
	return false
}
