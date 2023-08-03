package user

type (
	AccStatus int
)

const (
	AccNormal AccStatus = iota
	AccFreeze
	AccChangePwd
	AccStop
)

func (s AccStatus) IsValid() bool {
	switch s {
	case AccNormal, AccFreeze, AccChangePwd, AccStop:
		return true
	}
	return false
}

func (s AccStatus) Choices() string {
	switch s {
	case AccNormal:
		return "正常"
	case AccFreeze:
		return "冻结"
	case AccChangePwd:
		return "需要修改密码"
	case AccStop:
		return "需要修改密码"
	default:
		return "Unknown"
	}
}
