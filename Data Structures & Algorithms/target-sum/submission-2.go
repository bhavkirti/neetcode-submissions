func findTargetSumWays(nums []int, target int) int {
    dp := make(map[int]int)
	dp[0] = 1

	for _, num := range nums{
		next := make(map[int]int)
		for total, count := range dp{
			next[total+num] += count
			next[total-num] += count
		}	
		dp = next
	}
	return dp[target]
}
