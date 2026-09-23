func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
	if m < n {
		m, n = n, m
		word1, word2 = word2, word1
	}

	dp := make([]int, n+1)
	for j :=0; j <= n; j++ {
		dp[j] = n-j
	}

	for i := m-1; i >=0; i-- {
		next := dp[n]
		dp[n] = m-i
		for j := n-1; j >= 0; j-- {
			temp:=dp[j]
			if word1[i] == word2[j]{
				dp[j] = next
			} else{
				dp[j] = 1 + min(dp[j], min(dp[j+1], next))
			}
			next = temp
		}
	}
	return dp[0]
}

func min(a,b int)int{
	if a < b{
		return a
	}
	return b
}
