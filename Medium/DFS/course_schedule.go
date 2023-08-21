package DFS

func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int, numCourses)
	visited := make(map[int]int, numCourses)

	if len(prerequisites) == 0 {
		return true
	}

	for _, val := range prerequisites {
		preMap[val[0]] = append(preMap[val[0]], val[1])
	}

	var dfs func(course int) bool

	dfs = func(course int) bool {
		if visited[course] == 1 {
			return false
		}

		if len(preMap[course]) == 0 {
			return true
		}

		visited[course] = 1

		for _, val := range preMap[course] {
			if !dfs(val) {
				return false
			}
		}

		visited[course] = -1
		preMap[course] = []int{}
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
