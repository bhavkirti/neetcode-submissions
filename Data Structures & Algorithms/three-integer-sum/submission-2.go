func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := [][]int{}
	for i:=0; i < n; i++ {
		a := nums[i]

		if i > 0 && nums[i-1] == a{
			continue
		}

		l,r := i+1, len(nums)-1

		for l < r {
			sum := a+ nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{a,nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1]{
					l++
				}
			}
		}
	}
	return res
}
