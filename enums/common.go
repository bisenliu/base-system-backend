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

func (BoolSign) Choices(key BoolSign) string {
	boolChoices := map[BoolSign]string{
		False: "失败",
		True:  "成功",
	}
	desc, ok := boolChoices[key]
	if !ok {
		return "失败"
	}
	return desc
}
