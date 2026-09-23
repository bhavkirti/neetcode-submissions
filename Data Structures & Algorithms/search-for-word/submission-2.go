func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var dfs func(r int, c int, i int) bool 
	dfs = func(r int, c int, i int)bool{
		if i == len(word){
			return true
		}

		if r < 0 || c < 0|| r >= rows || c >= cols || board[r][c] != word[i] || board[r][c] == '#' {
			return false
		}

		temp := board[r][c]
		board[r][c] = '#'
		res := dfs(r+1, c, i+1) || dfs(r-1, c, i+1) || dfs(r, c+1, i+1) || dfs(r, c-1, i+1)
		board[r][c] = temp
		return res
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols ; j++{
			if dfs(i, j, 0){
				return true
			}
		}
	}
	return false
}
