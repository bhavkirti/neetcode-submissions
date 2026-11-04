func reverse(x int) int {
	var rev func(int, int)int
	rev = func(n, rec int) int {
		if n == 0{
			return rec
		}
		rec = rec*10 + n%10
		return rev(n/10, rec)
	}
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	reversedNum := rev(x, 0) * sign
	if reversedNum < -(1<<31) || reversedNum > (1<<31-1){
		return 0
	}
	return reversedNum
}
