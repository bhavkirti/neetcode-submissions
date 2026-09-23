type MinHeap []int

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i,j int) bool {return h[i] < h[j]}
func (h MinHeap) Swap(i,j int) {h[i], h[j] = h[j], h[i]}
func (h *MinHeap) Push(x interface{})  {*h = append(*h, x.(int))}
func (h *MinHeap) Pop() interface{}{
	old := *h
	n:= len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func lastStoneWeight(stones []int) int {
	pq := &MinHeap{}
	heap.Init(pq)

	for _, s := range stones{
		heap.Push(pq, -s)
	}

	for pq.Len() > 1 {
		first := heap.Pop(pq).(int)
		second := heap.Pop(pq).(int)
		if second > first {
			heap.Push(pq, first-second)
		}
	}
	heap.Push(pq, 0)
	return -heap.Pop(pq).(int)
}
