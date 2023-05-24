package enums

type BoolSign int

const (
	False BoolSign = iota
	True
)

func (receiver BoolSign) IsValid() bool {
	switch receiver {
	case False, True:
		return true
	}
	return false
}
