func reverseBits(n int) int {
    var res uint32 = 0
    ns := uint32(n) // convert to uint32 explicitly for 32-bit operations
    
    for i := 0; i < 32; i++ {
        bit := (ns >> i) & 1        // FIX: Shift by 'i', not '1'
        res |= (uint32(bit) << (31 - i))
    }
    
    return int(res)
}