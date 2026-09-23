func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a, b := nums1, nums2
	total := len(nums1) + len(nums2)
	half := (total+1)/2

	if len(a) > len(b){
		a, b = b, a
	}

	l, r := 0, len(a)

	for l <= r {
		i := (l+r)/2
		j := half-i

		aleft := math.MinInt64
		if i > 0 {
			aleft = a[i-1]
		}
		aright := math.MaxInt64
		if i < len(a){
			aright = a[i]
		}

		bleft := math.MinInt64
		if j > 0 {
			bleft = b[j-1]
		}
		bright := math.MaxInt64
		if j < len(b) {
			bright = b[j]
		}

		if aleft <= bright && bleft <= aright {
			if total%2 != 0 {
				return float64(max(aleft, bleft))
			}
			return (float64(max(aleft, bleft)) + float64(min(aright, bright)))/2.0
		} else if aleft > bright {
			r = i-1
		} else {
			l = i+1
		}
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
