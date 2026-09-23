func generateParenthesis(n int) []string {
	stack := make([]string, 0)
	res := make([]string, 0)

	var backtracking func(int, int)
	backtracking = func(open, closed int){
		if open == n && closed == n {
			res = append(res, strings.Join(stack, ""))
			return
		}

		if open < n {
			stack = append(stack, "(")
			backtracking(open+1, closed)
			stack = stack[:len(stack)-1]
		}

		if closed < open {
			stack = append(stack, ")")
			backtracking(open, closed+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtracking(0,0)
	return res
}
