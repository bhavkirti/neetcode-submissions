func productExceptSelf(nums []int) []int {
	n := len(nums)
    result := make([]int, n)

	for i:=0; i < n; i++ {
		result[i] = 1
	}

	prefix, suffix := 1, 1

	for i:=0; i < n; i++{
		result[i] = prefix
		prefix *= nums[i]
	}

	for i :=n-1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}
