func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	var backtrack func(int, []int)
	backtrack = func(i int, subset []int) {
		res = append(res, append([]int{}, subset...))

		for j := i; j < len(nums); j++ {
			if j > i && nums[j] == nums[j-1]{
				continue
			}
			subset = append(subset, nums[j])
			backtrack(j+1, subset)
			subset = subset[:len(subset)-1]
		}
	}
	
	backtrack(0, []int{})
	return res
}
