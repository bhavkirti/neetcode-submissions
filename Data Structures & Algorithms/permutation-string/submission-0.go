func checkInclusion(s1 string, s2 string) bool {
	s1Sorted := []rune(s1)
	sort.Slice(s1Sorted, func(i, j int) bool {
		return s1Sorted[i] < s1Sorted[j]
	})
	s1 = string(s1Sorted)

	for i := 0; i < len(s2); i++ {
			r := i+len(s1)
			if r > len(s2){
				return false
			}
			subStr := s2[i : r]
			subStrSorted := []rune(subStr)
			sort.Slice(subStrSorted, func(a, b int) bool {
				return subStrSorted[a] < subStrSorted[b]
			})
			if string(subStrSorted) == s1 {
				return true
			}
		
	}
	return false
}
