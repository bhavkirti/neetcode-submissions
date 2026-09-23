func countBits(n int) []int {
	res := make([]int, n+1)

	for i:=0; i <=n ; i++ {
		one:=0
		for j:=0;j < 32; j++ {
			if i&(1<<j) != 0{
				one++
			}
		}
		res[i] = one
	}
	return res
}
