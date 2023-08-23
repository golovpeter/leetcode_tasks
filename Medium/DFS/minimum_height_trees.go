package DFS

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make([][]int, n)
	degree := make([]int, n)

	for _, val := range edges {
		graph[val[0]] = append(graph[val[0]], val[1])
		graph[val[1]] = append(graph[val[1]], val[0])
		degree[val[0]]++
		degree[val[1]]++
	}

	var queue []int

	for i := range degree {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	for n > 2 {
		size := len(queue)
		n -= size

		for size > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, v := range graph[node] {
				degree[v]--

				if degree[v] == 1 {
					queue = append(queue, v)
				}
			}

			size--
		}
	}

	return queue
}
