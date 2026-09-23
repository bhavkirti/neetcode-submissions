func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		res[sortedStr] = append(res[sortedStr], str)
	}

	var result [][]string
	for _, group := range res {
		result = append(result, group)
	}
	return result
}

func sortString(str string) string {
	chars := []rune(str)
	sort.Slice(chars, func(i ,j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}
