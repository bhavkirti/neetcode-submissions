func maxProduct(nums []int) int {
    res := nums[0]
	curMax, curMin := 1, 1

	for _, num := range nums{
		if num < 0{
			curMax, curMin = curMin, curMax
		}

		curMax = max(num, num*curMax)
		curMin = min(num, num*curMin)

		res = max(res, curMax)
	}
	return res
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int{
	if a > b {
		return b
	}
	return a
}