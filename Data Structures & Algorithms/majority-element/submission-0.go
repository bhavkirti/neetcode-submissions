func majorityElement(nums []int) int {
    res, maxCount := 0,0 
	count := make(map[int]int)

	for _, num := range nums{
		count[num]++
		if count[num] > maxCount{
			res = num
			maxCount = count[num]
		}
	}
	return res
}
