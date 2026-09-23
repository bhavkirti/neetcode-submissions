func findMin(nums []int) int {
	res := nums[0]
	l, r := 0, len(nums)-1

	for l <= r {
		if nums[l] < nums[r] {
			res = min(res, nums[l])
			break
		}

		mid := (l+r)/2 
		res = min(res, nums[mid])

		if nums[mid] >= nums[l] {
			l = mid + 1
		} else{
			r = mid - 1
		}
	}
	return res
}


func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}
