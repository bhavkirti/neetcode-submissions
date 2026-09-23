func leastInterval(tasks []byte, n int) int {
	count := make(map[byte]int)
	for _, task := range tasks{
		count[task]++
	}

	maxHeap := &MaxHeap{}
	for _, cnt := range count {
		*maxHeap = append(*maxHeap, cnt)
	}
	heap.Init(maxHeap)

	time := 0
	q := make([][2]int, 0)

	for maxHeap.Len() > 0 || len(q)> 0{
		time++
		if maxHeap.Len() == 0{
			time = q[0][1]
		} else{
			cnt := heap.Pop(maxHeap).(int)-1
			if cnt > 0{
				q = append(q, [2]int{cnt, time+n})
			}
		}
		if len(q)>0 && q[0][1] == time{
			heap.Push(maxHeap, q[0][0])
			q = q[1:]
		}
	}
	return time
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}