type NumArray struct {
    prefix [] int
}


func Constructor(nums []int) NumArray {
    prefixarr := make([]int, len(nums)+1)
	prefixarr[0] = 0
	cur := 0
	for i , num := range nums{
		cur += num
		prefixarr[i+1] = cur
	}
	return NumArray{prefix:prefixarr}
}


func (this *NumArray) SumRange(left int, right int) int {
	l:= 0
    l = this.prefix[left]
	r := this.prefix[right+1]

	return r-l
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */