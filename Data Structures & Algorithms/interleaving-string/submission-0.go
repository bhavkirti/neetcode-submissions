func isInterleave(s1 string, s2 string, s3 string) bool {
    m, n := len(s1), len(s2)
	if m+n != len(s3){
		return false
	}
	if m > n{
		m, n = n, m
		s1, s2 = s2, s1
	}

	dp := make([]bool, n+1)
	dp[n] = true

	for i := m; i >=0; i--{
		next := (m==i)
		for j:=n; j>=0; j--{
			res := next
			if j < n {
				res = false
			}
			if i < m && s1[i]==s3[i+j] && dp[j]{
				res = true
			}
			if j < n && s2[j]==s3[i+j] && next {
				res = true
			}
			dp[j] = res
			next = dp[j]
		}
	}
	return dp[0]
}
