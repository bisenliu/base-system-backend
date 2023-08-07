package common

// RemoveDuplication
//  @Description: Slice 去重
//  @param arr Slice
//  @return []string 去重后的结果

func RemoveDuplication(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}
	return arr[:j]
}
