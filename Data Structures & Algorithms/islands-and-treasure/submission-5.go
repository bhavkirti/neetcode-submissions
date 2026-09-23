func islandsAndTreasure(grid [][]int) {
    rows, cols := len(grid), len(grid[0])
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    q := [][2]int{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				q = append(q, [2]int{i,j})
			}
		}
	}
	if len(q) == 0 {
		return
	}

    for len(q) > 0 {
		node:=q[0]
		q = q[1:]
		row, col := node[0], node[1]

		for _, dir := range directions {
			r, c := row+dir[0], col+dir[1]
			if r >= rows || c >= cols || r < 0 || c < 0 || grid[r][c] != 2147483647 {
				continue
			}
			q = append(q, [2]int{r, c})
			grid[r][c] = grid[row][col] +1
		}
	}
}