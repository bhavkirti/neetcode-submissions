func productExceptSelf(nums []int) []int {
    result := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        result[i] = 1
        for j:= 0; j < len(nums); j++ {
            if i != j {
                result[i] *= nums[j]
            }
        }
    }

    return result

}
