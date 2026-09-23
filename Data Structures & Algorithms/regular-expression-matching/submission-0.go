func isMatch(s string, p string) bool {
    m , n := len(s), len(p)
	dp := make([]bool, n+1)
	dp[n] = true

	for i := m ; i >= 0; i-- {
		dp1 := dp[n]
		dp[n] = (i == m)

		for j := n-1; j >= 0; j--{
			match := i < m && (s[i] == p[j] || p[j] == '.')
			res := false
			if j+1 < n && p[j+1] == '*' {
				res = dp[j+2]
				if match {
					res = res || dp[j]
				}
			} else if match {
				res = dp1
			}
			dp[j], dp1 = res, dp[j]
		}
	}
	return dp[0]
}
