func hasDuplicate(nums []int) bool {
    seen := make(map[int]struct{})

    for _, n:= range nums {
        seen[n] = struct{}{}
    }

    return len(seen) < len(nums)
}
