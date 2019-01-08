package MaximumSubarray

import (
	. "Algorithms/LeetCode/internal/box/math"
)

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	maxes := make([][]int, length)
	maxes[0] = []int{-1 << 31, nums[0]}
	for i := 1; i < length; i++ {
		maxes[i] = make([]int, 2)
		maxes[i][0] = MaxOfTwo(maxes[i-1][0], maxes[i-1][1])
		maxes[i][1] = MaxOfTwo(0, maxes[i-1][1]) + nums[i]
	}
	return MaxOfTwo(maxes[length-1][0], maxes[length-1][1])
}
