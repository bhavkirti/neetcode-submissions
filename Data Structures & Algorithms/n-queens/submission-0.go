func solveNQueens(n int) [][]string {
	col := make(map[int]bool)
	posDiag := make(map[int]bool)
	negDiag := make(map[int]bool)
	board := make([][]rune, n)
	var res [][]string
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i]{
			board[i][j] = '.'
		}
	}

	var backtrack func(r int)
	backtrack = func (r int){
		if r == n {
			sol := make([]string, n)
			for i := range board{
				sol[i] = string(board[i])
			}
			res = append(res, sol)
			return
		}

		for c := 0; c < n; c++ {
			if col[c] || posDiag[c+r] || negDiag[r-c]{
				continue
			}

			col[c] = true
			posDiag[c+r] = true
			negDiag[r-c] = true
			board[r][c] = 'Q'

			backtrack(r+1)

			col[c] = false
			posDiag[c+r] = false
			negDiag[r-c] = false
			board[r][c] = '.'
		}
	}
	backtrack(0)
	return res
}
