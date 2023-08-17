package BFS

type Node struct {
	Row    int
	Column int
}

var direction = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func updateMatrix(mat [][]int) [][]int {
	n, m := len(mat), len(mat[0])
	queue := make([]Node, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, Node{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, dir := range direction {
			x, y := node.Row+dir[0], node.Column+dir[1]

			if (x >= 0) && (x < len(mat)) && (y >= 0) && (y < len(mat[0])) && mat[x][y] == -1 {
				mat[x][y] = mat[node.Row][node.Column] + 1
				queue = append(queue, Node{x, y})
			}
		}
	}

	return mat
}
