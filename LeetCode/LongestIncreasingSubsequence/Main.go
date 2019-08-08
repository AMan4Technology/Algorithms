package LongestIncreasingSubsequence

import (
	"Algorithms"
)

func lengthOfLIS(nums []int) (max int) {
	length := len(nums)
	if length < 2 {
		return length
	}
	lengths := make([]int, length)
	for i, num := range nums {
		for j := i - 1; j >= 0; j-- {
			if nums[j] < num {
				lengths[i] = Algorithms.MaxOfTwo(lengths[j]+1, lengths[i])
			}
		}
		max = Algorithms.MaxOfTwo(lengths[i], max)
	}
	return
}
