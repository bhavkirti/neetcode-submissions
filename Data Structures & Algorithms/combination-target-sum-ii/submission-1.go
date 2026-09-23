func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)

	var dfs func(int, []int, int)

	dfs = func(idx int, path []int, cur int){
		if cur == target{
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := idx ; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1]{
				continue
			}
			if cur + candidates[i] > target {
				break
			}

			path = append(path, candidates[i])
			dfs(i+1, path, cur + candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}
