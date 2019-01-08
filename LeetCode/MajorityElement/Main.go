package MajorityElement

import "math"

func majorityElement(nums []int) int {
	length := len(nums)
	if length == 0 {
		return math.MinInt32
	}
	var (
		result int
		count  int
	)
	for _, val := range nums {
		if val == result {
			count++
		} else if count == 0 {
			result, count = val, 1
		} else {
			count--
		}
	}
	count = 0
	for _, val := range nums {
		if val == result {
			count++
		}
	}
	if count > length/2 {
		return result
	}
	return math.MinInt32
}
