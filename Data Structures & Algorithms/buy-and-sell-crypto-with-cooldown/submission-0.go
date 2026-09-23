func maxProfit(prices []int) int {
    n := len(prices)
	dp1B, dp1S, dp2B := 0,0,0

	for i := n-1; i >=0; i--{
		dpB := max(dp1S - prices[i], dp1B)
		dpS := max(dp2B + prices[i], dp1S)
		dp2B, dp1S = dp1B, dpS
		dp1B = dpB
	}
	return dp1B
}

func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}
