func maxProfit(prices []int) int {
	maxProfit := 0
	minBuy := math.MaxInt32

	for _, sell := range prices {
		if sell - minBuy > maxProfit {
			maxProfit = sell - minBuy
		} 
		if sell < minBuy {
			minBuy = sell
		}
	}
	return maxProfit
}
