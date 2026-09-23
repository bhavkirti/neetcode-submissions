func maxProfit(prices []int) int {
    maxProfit := 0
    l,r := 0,1

    for r < len(prices) {
        if prices[r] > prices[l] {
            profit := prices[r]-prices[l] 
            if profit > maxProfit {
                maxProfit = profit
            }
        } else{
            l=r
        }
        r++
    }
    return maxProfit
}
