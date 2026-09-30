func maxArea(heights []int) int {
	maxA := 0
	l, r := 0, len(heights)-1

	for l < r {
		area := min(heights[l], heights[r]) * (r-l)

		if heights[l] < heights[r]{
			l++
		} else{
			r--
		}
		maxA = max(area, maxA)
	}
	return maxA
}

func min(a, b int) int{
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
