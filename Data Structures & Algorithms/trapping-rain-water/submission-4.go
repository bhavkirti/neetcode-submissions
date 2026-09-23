func trap(height []int) int {
	n := len(height)
	sum := 0
	if n == 0{
		return 0
	}
	l, r := 0, n-1
	lMax, rMax := height[l], height[r]
	
	for l < r {
		if lMax < rMax {
			l++
			lMax = max(lMax, height[l])
			sum += lMax - height[l]
		} else {
			r--
			rMax = max(rMax, height[r])
			sum += rMax - height[r]
		}
	}
	return sum
}