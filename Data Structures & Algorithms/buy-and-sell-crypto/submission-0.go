func maxProfit(prices []int) int {
    maxProfit := 0
    diff := 0
    n := len(prices)
    for i := 0; i < n; i++ {
        for j:= i+1; j < n; j++ {
            diff = prices[j]-prices[i] 
            if diff > maxProfit {
                maxProfit = diff
            }
        }
    }
    return maxProfit
}
