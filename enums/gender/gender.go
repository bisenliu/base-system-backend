package gender

type Gender int

const (
	Female Gender = iota // 女
	Male                 // 男
)

func (receiver Gender) IsValid() bool {
	switch receiver {
	case Female, Male:
		return true
	}
	return false
}
