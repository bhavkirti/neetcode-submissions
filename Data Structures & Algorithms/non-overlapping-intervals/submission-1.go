func eraseOverlapIntervals(intervals [][]int) int {
    res := 0
	sort.Slice(intervals, func(i,j int)bool {
		return intervals[i][0] < intervals[j][0]
	})

	prevEnd := intervals[0][1]

	for _, interval := range intervals[1:] {
		if interval[0] >= prevEnd {
			prevEnd = interval[1]
		} else {
			res++ 
			prevEnd = min(interval[1], prevEnd)
		}
	}
	return res
}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}
