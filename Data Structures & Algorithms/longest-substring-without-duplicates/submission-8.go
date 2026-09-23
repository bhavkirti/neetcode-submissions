func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)

	l, res := 0, 0

	for i := 0; i < len(s); i++ {
		if idk, found := mp[s[i]]; found {
			l = max(idk+1, l)
		}
		mp[s[i]] = i
		if i-l+1 > res {
			res = i-l+1
		}
	}
	return res
}
