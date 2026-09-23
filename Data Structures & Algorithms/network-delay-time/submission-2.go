type Edge struct {
	weight, node int
}
 
type MinHeap []Edge

func(h MinHeap) Len() int { return len(h)}
func(h MinHeap) Less(i, j int) bool {return h[i].weight < h[j].weight}
func(h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func(h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Edge))}
func(h *MinHeap) Pop() interface{}{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
    edges := make(map[int][]Edge)

	for _, time := range times {
		u, v, w := time[0], time[1], time[2]
		edges[u] = append(edges[u], Edge{node: v, weight: w})
	}

	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, Edge{node: k , weight: 0})

	visited := make(map[int]bool)
	t := 0

	for pq.Len() > 0 {
		edge := heap.Pop(pq).(Edge)
		node, time := edge.node, edge.weight

		if visited[node] {
			continue
		}
		visited[node] = true
		t = time

		for _, next := range edges[node]{
			if !visited[next.node]{
				heap.Push(pq, Edge{node: next.node, weight: time + next.weight})
			}
		}
	}
	if len(visited) == n {
		return t
	}
	return -1
}
