package LongestIncreasingSubsequence

import . "Algorithms/LeetCode/internal/box/math"

func lengthOfLIS(nums []int) (max int) {
	length := len(nums)
	if length < 2 {
		return length
	}
	lengths := make([]int, length)
	for i, num := range nums {
		for j := i - 1; j >= 0; j-- {
			if nums[j] < num {
				lengths[i] = MaxOfTwo(lengths[j]+1, lengths[i])
			}
		}
		max = MaxOfTwo(lengths[i], max)
	}
	return
}
