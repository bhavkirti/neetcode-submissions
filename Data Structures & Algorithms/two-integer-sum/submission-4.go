func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, n := range nums {
		diff := target - nums[i]
		if j, found := hashMap[diff]; found {
			return []int{j, i}
		}
		hashMap[n] = i
	}
	return []int{}
}
