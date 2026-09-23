func longestCommonSubsequence(text1 string, text2 string) int {
    if len(text1) < len(text2){
		text1, text2 = text2, text1
	}

	dp := make([]int, len(text2)+1)

	for i := len(text1)-1; i >=0; i--{
		prev := 0
		for j := len(text2)-1; j >= 0; j--{
			temp := dp[j]
			if text1[i] == text2[j]{
				dp[j] = 1+prev
			} else{
				dp[j] = max(dp[j], dp[j+1])
			}
			prev = temp
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}