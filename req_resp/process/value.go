package process

// 判断值是否为-10000
func ValueJudge(array *[]int32)(isGoodValue bool) {
	isGoodValue = true
	for i:=0; i < len(*array); i++ {
		if (*array)[i] == -10000 {
			return false
		}
	}
	return true
	}