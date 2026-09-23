type Pair struct{
	row, col int
}
func orangesRotting(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
	fresh := 0
	time := 0
	queue := make([]Pair, 0)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++{
			if grid[r][c] == 1{
				fresh++
			}
			if grid[r][c] == 2 {
				queue = append(queue, Pair{r,c})
			}
		}
	}

	directions := [][]int{{-1,0}, {1, 0}, {0,-1}, {0,1}}

	for fresh > 0 && len(queue) > 0{
		length := len(queue)
		for i := 0; i < length; i++ {
			current := queue[0]
			queue = queue[1:]

			for _, dir := range directions{
				newr := current.row + dir[0]
				newc := current.col + dir[1]
				
				if newr >=0 && newr < rows && newc >= 0 && newc < cols && grid[newr][newc] == 1 {
					grid[newr][newc] = 2
					queue = append(queue, Pair{newr, newc})
					fresh--
				}
			}
		}
		time++
	}
	if fresh == 0 {
		return time
	}
	return -1
}
