package utils

// 检查数组中是否有重复元素
func CheckRepeat[T int | string](s []T) bool {
	var slices []T
	for _, i := range s {
		if CheckContainsInSlice(slices, i) {
			return true
		} else {
			slices = append(slices, i)
		}
	}
	return false
}

func CheckContainsInSlice[T int | string](items []T, item T) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// IntersectArray 求两个切片的交集
func IntersectArray[T int | string](a []T, b []T) []T {
	var inter []T
	mp := make(map[T]bool)

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}
	return inter
}

// DiffArray 求两个切片的差集
func DiffArray[T int | string](a []T, b []T) []T {
	var diffArray []T
	temp := map[T]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}
	return diffArray
}
