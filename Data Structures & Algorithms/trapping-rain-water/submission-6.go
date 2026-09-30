func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	sum := 0
	l, r := 0, n-1
	lMax, rMax := height[l], height[r]

	for l < r {
		if lMax < rMax{
			l++
			lMax = max(height[l], lMax)
			sum += lMax - height[l]
		} else{
			r--
			rMax = max(height[r], rMax)
			sum += rMax - height[r]
		}
	}
	return sum
}
