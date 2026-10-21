func timeRequiredToBuy(tickets []int, k int) int {
    res := 0

	for i:=0; i < len(tickets); i++ {
		if i <= k {
			res += min(tickets[i], tickets[k])
		} else {
			res +=  min(tickets[i], tickets[k]-1)
		}
	}
	return res
}

func min(a, b int)int {
	if a < b {
		return a
	}
	return b
}