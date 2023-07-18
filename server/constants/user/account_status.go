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

func (receiver AccStatus) IsValid() bool {
	switch receiver {
	case AccNormal, AccFreeze, AccChangePwd, AccStop:
		return true
	}
	return false
}

func (receiver AccStatus) AccStatusDisplay(status AccStatus) string {
	switch status {
	case AccNormal:
		return "正常"
	case AccFreeze:
		return "冻结"
	case AccChangePwd:
		return "需要修改密码"
	default:
		return "停用"
	}
}
