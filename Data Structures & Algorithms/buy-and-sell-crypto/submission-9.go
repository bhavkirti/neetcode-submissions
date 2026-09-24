func maxProfit(prices []int) int {
	maxProfit := 0
	n := len(prices)
	l,r := 0, 1

	for r < n {
		if prices[r] > prices[l]{
			maxProfit = max(maxProfit, prices[r]-prices[l])
		} else {
			l = r
		}
		r++
	}
	return maxProfit
}

func mac(a, b int)int{
	if a > b {
		return a
	}
	return b
}
