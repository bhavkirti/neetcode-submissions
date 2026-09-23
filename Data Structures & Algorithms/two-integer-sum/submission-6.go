func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		req := target - nums[i]
		if j, found := hashMap[req] ; found {
			if i > j {
				return []int{j,i}
			}
			return []int{i,j}
		}
		hashMap[nums[i]] = i
	}
    return []int{-1,-1}
}
