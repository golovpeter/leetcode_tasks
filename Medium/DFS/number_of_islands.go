package DFS

func visit(grid [][]byte, i, j int) {
	if (i < 0) || (i > len(grid)-1) || (j < 0) || (j > len(grid[0])-1) || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	visit(grid, i-1, j)
	visit(grid, i+1, j)
	visit(grid, i, j-1)
	visit(grid, i, j+1)
}

func numIslands(grid [][]byte) int {
	res := 0
	n, m := len(grid), len(grid[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				visit(grid, i, j)
				res++
			}
		}
	}

	return res
}
