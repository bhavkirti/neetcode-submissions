func scoreOfString(s string) int {
	res := 0
	for i := 0; i < len(s)-1; i++ {
		diff := int(s[i])- int(s[i+1])
		if diff < 0 {
			diff = -diff
		}
		res += diff
	}
	return res
}
