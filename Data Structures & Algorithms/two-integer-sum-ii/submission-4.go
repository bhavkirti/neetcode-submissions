func twoSum(numbers []int, target int) []int {
	j:=len(numbers)-1
	i:= 0
	for i < len(numbers) && i < j{
			sum := numbers[i] + numbers[j] 
			if sum > target {
				j--
			} else if sum < target {
				i++
			} else {
				return []int{i+1, j+1}
			}
		
	}
	return []int{}
}
