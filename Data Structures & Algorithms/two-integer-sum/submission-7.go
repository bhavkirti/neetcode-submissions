func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		req := target - nums[i]
		if j, found := hashMap[req]; found {
			if i < j{
				return []int{i,j}
			}
			return []int{j,i}
		}
		hashMap[nums[i]] = i
	}
	return []int{-1,-1}
}
