package process

import (
	"fmt"
	"strconv"
)

func compareDate (a uint32, b uint32) (result bool) {
	strA := strconv.FormatUint(uint64(a),10)
	strB := strconv.FormatUint(uint64(b),10)
	// 如果A的年份比B大， 直接返回错误
	if strA[0:4] > strB[0:4] {
		return false
	// 如果A的年份比B小
	} else if strA[0:4] < strB[0:4]{
		return true
	// 如果年份相等, 比较月份
	} else {
		if strA[4:6] > strB[4:6] {
			return false
		} else if strA[4:6] < strB[4:6] {
			return true
		// 如果月份也相等, 当A日期大于等于B日期时返回错误
		} else {
			if strA[6:8] >= strB[6:8] {
				return false
			}
			return true
		}
	}
}

func DateAscend(array *[]uint32) (isAscend bool) {
	isAscend = true
	for i:=0; i < len(*array) - 1; i++ {
		if compareDate((*array)[i], (*array)[i+1]) == false {
			fmt.Printf("出问题的日期是：%d, %d",(*array)[i], (*array)[i+1])
			isAscend = false
			break
		}
	}
	return isAscend
}