func reverseBits(n int) int {
	var res uint32 = 0
	ns := uint32(n)
	for i := 0; i < 32; i++ {
		bit := (ns>>i)&1
		res |= (uint32(bit) << (31-i))
	}
	return int(res)
}

