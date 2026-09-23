func maxProfit(prices []int) int {
	maxProfit := 0

	l, r := 0, 1

	for r < len(prices) {
		if prices[r] > prices[l] {
			maxProfit = max(maxProfit, (prices[r]-prices[l]))
		} else{
			l = r
		}
		r++
	}
	return maxProfit
}

func max(a,b int) int{
	if a > b{
		return a
	}
	return b
}