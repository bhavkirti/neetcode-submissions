func maxArea(heights []int) int {
	l,r := 0, len(heights)-1
	maxArea := 0
	for l<r {
		area := minHeight(heights[l], heights[r]) * (r-l)
		if heights[r] > heights[l] {
			l++
		}else {
			r--
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func minHeight(height1, height2 int) int {
	if height1 > height2 {
		return height2
	}
	return height1
}
