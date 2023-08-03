package common

// In
//  @Description: 检测值是否在 Slice 里
//  @param target 值
//  @param strArray Slice
//  @return bool true/false

func In(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element {
			return true
		}
	}
	return false
}
