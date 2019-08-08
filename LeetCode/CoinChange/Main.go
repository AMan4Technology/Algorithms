package CoinChange

import (
	"Algorithms"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	nums := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		nums[i] = amount + 1
		for _, value := range coins {
			if value <= i {
				nums[i] = Algorithms.MinOfTwo(nums[i-value]+1, nums[i])
			}
		}
	}
	if nums[amount] > amount {
		return -1
	}
	return nums[amount]
}
