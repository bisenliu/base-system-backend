package gender

type Gender int

const (
	Female Gender = iota // 女
	Male                 // 男
)

func (g Gender) IsValid() bool {
	switch g {
	case Female, Male:
		return true
	}
	return false
}
