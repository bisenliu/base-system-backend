package user

type ChangePwdType int

// 修改密码类型
const (
	PwdChange ChangePwdType = iota // 密码修改
	SmsChange                      // 手机验证码修改
)

// IsValid
//  @Description: 修改密码类型枚举校验,配合自定义 validator
//  @receiver p 接收者
//  @return bool 是否校验通过

func (p ChangePwdType) IsValid() bool {
	switch p {
	case PwdChange, SmsChange:
		return true
	}
	return false
}
