package user

type ChangePwdType int

const (
	PwdChange ChangePwdType = iota // 密码修改
	SmsChange                      // 手机验证码修改
)

func (p ChangePwdType) IsValid() bool {
	switch p {
	case PwdChange, SmsChange:
		return true
	}
	return false
}
