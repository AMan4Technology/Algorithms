package MaximumProductSubarray

import (
	"Algorithms"
)

func maxProduct(nums []int) (max int) {
	length := len(nums)
	if length == 0 {
		return 0
	}
	var (
		maximums      = make([]int, 2)
		minusMaximums = make([]int, 2)
	)
	maximums[0], minusMaximums[0], max = nums[0], nums[0], nums[0]
	for i := 1; i < length; i++ {
		x, y := (i-1)%2, i%2
		maximums[y], minusMaximums[y] = Algorithms.MaxAndMin(maximums[x]*nums[i], minusMaximums[x]*nums[i], nums[i])
		max = Algorithms.MaxOfTwo(maximums[y], max)
	}
	return
}
