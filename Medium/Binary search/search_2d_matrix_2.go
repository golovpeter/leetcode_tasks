package Binary_search

func searchMatrix2(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1

	for i < n && j >= 0 {
		v := matrix[i][j]

		if v == target {
			return true
		} else if v > target {
			j--
		} else {
			i++
		}
	}

	return false
}
