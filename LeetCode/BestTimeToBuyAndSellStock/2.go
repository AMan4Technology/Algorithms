package BestTimeToBuyAndSellStock

func maxProfit2(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	var profit int
	for i := 1; i < length; i++ {
		if prices[i-1] < prices[i] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
