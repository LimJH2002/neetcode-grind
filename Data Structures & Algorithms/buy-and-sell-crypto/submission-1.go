func maxProfit(prices []int) int {
	futureMax := 0
	profit := 0

	for i := len(prices) - 1; i >= 0; i-- {
		futureMax = max(futureMax, prices[i])
		profit = max(profit, futureMax - prices[i])
	}

	return profit
}
