func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	prefix := make([]int, n)
	suffix := make([]int, n)
	prefix[0], suffix[n-1] = 1, 1
	for i:= 1; i < n; i++ {
		prefix[i] = nums[i-1]*prefix[i-1] 
	}

	for i:= n-2 ; i >=0 ; i-- {
		suffix[i] = nums[i+1]*suffix[i+1]
	}

	for i:=0; i < n ; i++ {
		result[i] = prefix[i]* suffix[i]
	}
	return result
}
