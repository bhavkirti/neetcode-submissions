func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	
	result := [][]int{}
	
	for i := 0; i < n; i++ {
		l, r := i+1, n-1
		if i != 0 && nums[i] == nums[i-1]{
			continue
		}


		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum > 0 {
				r--
			} else if sum < 0{
				l++
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				r--
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return result
}
