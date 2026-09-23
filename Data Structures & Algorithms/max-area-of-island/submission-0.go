func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0

	rows, columns := len(grid), len(grid[0])
	visit := make(map[[2]int]bool)

	var dfs func(r, c int) int

	dfs = func(r, c int) int{
		if r < 0 || c < 0 || r >= rows || c >= columns || grid[r][c] == 0 || visit[[2]int{r,c}]{
			return 0
		}

		visit[[2]int{r, c}] = true

		return 1+dfs(r+1,c) + dfs(r-1,c) + dfs(r, c+1) + dfs(r, c-1)
	}
	for r := 0; r < rows; r++{
		for c := 0; c < columns; c++{
			maxArea = max(maxArea, dfs(r, c))
		}
	}
	return maxArea
}

func max(a, b int)int{
	if a > b {
		return a
	}
	return b
}
