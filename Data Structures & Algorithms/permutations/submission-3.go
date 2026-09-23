func permute(nums []int) [][]int {
	res := [][]int{}
	backtrack(&res, []int{}, nums, make([]bool, len(nums)))
	return res
}

func backtrack(res *[][]int, perm []int, nums[]int, pick []bool){
	if len(perm) == len(nums) {
		temp := append([]int{}, perm...)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++{
		if !pick[i]{
			perm = append(perm, nums[i])
			pick[i] = true
			backtrack(res, perm, nums, pick)
			perm = perm[:len(perm)-1]
			pick[i] = false
		}
	}
}
