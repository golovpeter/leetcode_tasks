package DFS

func findOrder(numCourses int, prerequisites [][]int) []int {
	var res []int
	neighbor := make(map[int][]int)

	for _, val := range prerequisites {
		neighbor[val[0]] = append(neighbor[val[0]], val[1])
	}

	var dfs func(crs int) bool
	visit, cycle := make(map[int]bool), make(map[int]bool)

	dfs = func(crs int) bool {
		if cycle[crs] {
			return false
		}

		if visit[crs] {
			return true
		}

		cycle[crs] = true
		for _, val := range neighbor[crs] {
			if !dfs(val) {
				return false
			}
		}

		cycle[crs] = false
		visit[crs] = true
		res = append(res, crs)
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}
	}

	return res
}
