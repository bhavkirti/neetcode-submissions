func findDuplicate(nums []int) int {
    set := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := set[num]; exists{
			return num
		}
		set[num] = struct{}{}
	}

	return -1
}
