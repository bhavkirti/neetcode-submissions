func groupAnagrams(strs []string) [][]string {
	hashTable := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, s := range str {
			count[s-'a']++
		}
		hashTable[count] = append(hashTable[count], str)
	}
	var result [][]string
	for _, group := range hashTable{
		result = append(result, group)
	}
	return result
}
