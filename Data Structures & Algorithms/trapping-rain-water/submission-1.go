func trap(height []int) int {
	n := len(height) 
	if n == 0 {
		return 0
	}
	l,r := 0, n-1
	lMax, rMax := height[l], height[r]
	result := 0

	for l < r {
		if lMax < rMax {
			l++
			lMax = max(lMax, height[l])
			result += lMax - height[l]
		} else {
			r--
			rMax = max(rMax, height[r])
			result += rMax - height[r]
		}
	}
	return result
}

func max(h1, h2 int)int{
	if h1 > h2 {
		return h1
	}
	return h2
}
