func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	lastIndex := strings.LastIndex(s, " ")
	return len(s) - lastIndex-1
	}
