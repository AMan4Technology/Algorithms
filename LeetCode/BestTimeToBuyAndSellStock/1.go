package BestTimeToBuyAndSellStock

import (
	"Algorithms"
)

func maxProfit1(prices []int) (max int) {
	if len(prices) < 2 {
		return 0
	}
	var buyDay int
	for day, price := range prices {
		priceChange := price - prices[buyDay]
		if priceChange > 0 {
			max = Algorithms.MaxOfTwo(priceChange, max)
		} else if priceChange < 0 {
			buyDay = day
		}
	}
	return
}
