func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

	max := 0
	curr, streak := nums[0], 0
	i:=0
	for i < len(nums){
		if curr != nums[i]{
			curr = nums[i]
			streak = 0
		}
		for i < len(nums) && nums[i] == curr {
			i++
		}
		streak++
		curr++
		if streak > max {
			max = streak
		}
	}
	return max
}
