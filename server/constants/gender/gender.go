package gender

type Gender int

// 性别枚举
const (
	Female Gender = iota // 女
	Male                 // 男
)

// IsValid
//  @Description: 性别枚举校验,配合自定义 validator
//  @receiver g 接收者
//  @return bool 是否校验通过

func (g Gender) IsValid() bool {
	switch g {
	case Female, Male:
		return true
	}
	return false
}
