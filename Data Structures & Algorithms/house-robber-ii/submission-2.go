func rob(nums []int) int {
    return max(nums[0], max(helper(nums[1:]), helper(nums[:len(nums)-1])))
}

func helper(nums []int)int{
	rob1, rob2 := 0, 0
	for _, num := range nums{
		temp := max(num+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
