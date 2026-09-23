func singleNumber(nums []int) int {
	seen := make(map[int]bool)

	for _, num := range nums{
		if !seen[num]{
			seen[num] = true
		} else{
			delete(seen, num)
		}
	}

	for num := range seen {
		return num
	}
	return -1
}
