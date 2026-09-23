func missingNumber(nums []int) int {
	n := len(nums)
	xorr := n
	for i := 0; i < n; i++{
		xorr = xorr^i^nums[i]
	}
	return xorr
}
