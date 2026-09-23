func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([][2]int, 0)

	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			maxArea = max(maxArea, top[1]*(i-top[0]))
			start = top[0]
		}
		stack = append(stack, [2]int{start, h})
	}
	for _, h := range stack {
		maxArea = max(maxArea, h[1]*(len(heights)-h[0]))
	}
	return maxArea
}

func max(a,b int) int {
	if a > b{
		return a
	}
	return b
}