func findOrder(numCourses int, prerequisites [][]int) []int {
    preMap := make(map[int][]int)
	for i := 0; i < numCourses ; i++ {
		preMap[i] = []int{}
	}

	for _, prereq := range prerequisites{
		crs, pre := prereq[0], prereq[1]
		preMap[crs] = append(preMap[crs], pre)
	}
	output:= []int{}
	visit := make(map[int]bool)
	cycle := make(map[int]bool)

	var dfs func(int)bool
	dfs = func(crs int) bool {
		if visit[crs]{
			return true
		}
		if cycle[crs]{
			return false
		}
		cycle[crs] = true
		
		for _, pre := range preMap[crs]{
			if !dfs(pre) {
				return false
			}
		}
		visit[crs] = true
		cycle[crs] = false
		output = append(output, crs)
		return true
	}

	for c := 0; c < numCourses; c++{
		if !dfs(c){
			return []int{}
		}
	}
	return output
}
