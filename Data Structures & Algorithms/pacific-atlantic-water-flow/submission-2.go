func pacificAtlantic(heights [][]int) [][]int {
	rows, columns := len(heights), len(heights[0])

	pcf := make(map[[2]int] bool)
	atl := make(map[[2]int] bool)

	
	var dfs func(r, c int, visit map[[2]int]bool, prevHeight int)
    dfs = func(r, c int, visit map[[2]int]bool, prevHeight int){
		coord := [2]int{r,c}
		if visit[coord] || r < 0 || c < 0 || r == rows || c == columns || heights[r][c] < prevHeight {
			return
		}
		visit[coord] = true
		dfs(r+1, c, visit, heights[r][c])
		dfs(r-1, c, visit, heights[r][c])
		dfs(r, c+1, visit, heights[r][c])
		dfs(r, c-1, visit, heights[r][c])
	}

	for c := 0; c < columns; c++ {
		dfs (0, c, pcf, heights[0][c])
		dfs (rows-1, c, atl, heights[rows-1][c])
	}

	for r := 0; r < rows; r++ {
		dfs (r, 0, pcf, heights[r][0])
		dfs (r, columns-1, atl, heights[r][columns-1])
	}
	res := make([][]int, 0)
	for r:=0; r<rows; r++{
		for c:=0; c<columns; c++{
			coord := [2]int{r,c}
			if pcf[coord] && atl[coord]{
				res = append(res, []int{r,c})
			}
		}
	}
	return res
}
