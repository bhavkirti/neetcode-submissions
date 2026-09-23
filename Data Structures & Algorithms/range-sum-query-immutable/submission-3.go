type NumArray struct {
    prefix []int
}


func Constructor(nums []int) NumArray {
    prefixArr := make([]int, len(nums))
	cur := 0
	for i, num := range nums {
		cur += num
		prefixArr[i] = cur
	}
	return NumArray{prefix: prefixArr}
}


func (this *NumArray) SumRange(left int, right int) int {
    
	l := 0
	r := this.prefix[right]
	if left > 0 {
		l = this.prefix[left-1]
	} 
	return r-l
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */