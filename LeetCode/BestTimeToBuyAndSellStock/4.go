package BestTimeToBuyAndSellStock

import . "Algorithms/LeetCode/internal/box/math"

func maxProfit(prices []int, transNum int) (max int) {
	length := len(prices)
	if length <= 1 || transNum == 0 {
		return 0
	}
	if transNum >= length/2 {
		for i := 1; i < length; i++ {
			if prices[i-1] < prices[i] {
				max += prices[i] - prices[i-1]
			}
		}
		return
	}
	maximums := make([][][]int, 2)
	for day := range maximums {
		maximums[day] = make([][]int, 2)
		for num := range maximums[day] {
			maximums[day][num] = make([]int, transNum+1)
		}
	}
	for trans := range maximums[0][1] {
		maximums[0][1][trans] = -prices[0]
	}
	for day := 1; day < length; day++ {
		x, y := (day-1)%2, day%2
		maximums[y][1][0] = MaxOfTwo(maximums[x][1][0], maximums[x][0][0]-prices[day])
		for trans := 1; trans <= transNum; trans++ {
			maximums[y][0][trans] = MaxOfTwo(maximums[x][0][trans], maximums[x][1][trans-1]+prices[day])
			maximums[y][1][trans] = MaxOfTwo(maximums[x][1][trans], maximums[x][0][trans]-prices[day])
		}
	}
	return Max(0, maximums[(length-1)%2][0]...)
}
