func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	freq := make([][]int, len(nums)+1)
	res := []int{}

	for _, num := range nums{
		hashMap[num]++
	}

	for num, cnt := range hashMap{
		freq[cnt] = append(freq[cnt], num)
	}

	for i := len(freq)-1; i >= 0; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res)==k{
				return res
			}
		}
	}
	return res
}
