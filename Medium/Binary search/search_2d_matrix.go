package Binary_search

func searchMatrix(matrix [][]int, target int) bool {
	rows, columns := len(matrix), len(matrix[0])

	top, bot := 0, rows-1
	for top <= bot {
		mid := (top + bot) / 2

		if target > matrix[mid][len(matrix[0])-1] {
			top = mid + 1
		} else if target < matrix[mid][0] {
			bot = mid - 1
		} else {
			break
		}
	}

	if !(top <= bot) {
		return false
	}

	l, r := 0, columns-1
	row := (top + bot) / 2

	for l <= r {
		mid := (l + r) / 2

		if matrix[row][mid] < target {
			l = mid + 1
		} else if matrix[row][mid] > target {
			r = mid - 1
		} else {
			return true
		}
	}

	return false
}
