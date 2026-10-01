func characterReplacement(s string, k int) int {
	mp := make(map[byte]int)
	res, l, maxf := 0,0,0
	for r := 0; r < len(s); r++ {
		mp[s[r]]++
		if mp[s[r]] > maxf {
			maxf = mp[s[r]]
		}

		for (r-l+1) - maxf > k {
			mp[s[l]]--
			l++
		}
		
		if r-l+1 > res {
			res = r-l+1
		}
	}
	return res
}
