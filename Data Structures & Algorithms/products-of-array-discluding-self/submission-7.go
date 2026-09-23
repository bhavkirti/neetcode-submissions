func productExceptSelf(nums []int) []int {
	totalProduct := 1
	zeroCount := 0

	for _, v := range nums {
		if v != 0 {
			totalProduct *= v
		} else {
			zeroCount++
		}
	}

	result := make([]int, len(nums))

	if zeroCount > 1 {
		return result
	}

	for i, v := range nums {
		if zeroCount > 0 {
			if v == 0 {
				result[i] = totalProduct
			} else{
				result[i] = 0
			}
		} else {
			result[i] = totalProduct/v
		}
	}
	return result
}
