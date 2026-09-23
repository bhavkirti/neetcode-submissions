func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)

	sort.Ints(merged)
	n := len(merged)

	if n%2 == 0 {
		m1 := n/2
		m2 := (n - 1)/2
		return float64(merged[m1]+merged[m2])/2.0
	} else{
		return float64(merged[n/2])
	}
}
