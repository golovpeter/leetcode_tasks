package Arrays

func spiralOrder(matrix [][]int) []int {
	var res []int

	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for {
		//right
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}

		top++
		if top > bottom {
			break
		}

		//down
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}

		right--
		if left > right {
			break
		}

		//left
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}

		bottom--
		if top > bottom {
			break
		}

		//up
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}

		left++
		if left > right {
			break
		}
	}

	return res
}
