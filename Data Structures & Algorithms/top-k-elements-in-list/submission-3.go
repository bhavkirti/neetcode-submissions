func topKFrequent(nums []int, k int) []int {
	hashTable := make(map[int]int)

	for _, n := range nums {
		hashTable[n]++
	}

	arr := make([][2]int, 0, len(hashTable))

	for num, count := range hashTable{
		arr = append(arr, [2]int{count, num})
	}

	sort.Slice(arr, func(i,j int) bool{
		return arr[i][0] > arr[j][0]
	})

	res := make([]int, k)

	for i:=0; i < k; i++ {
		res[i] = arr[i][1]
	}
	return res
}
