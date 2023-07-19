package Medium

func generateParenthesis(n int) []string {
	var res []string
	var backtrack func(left, right int, curr []byte)

	backtrack = func(left, right int, curr []byte) {
		if left+right == n*2 {
			res = append(res, string(curr))
			return
		}

		if left < n {
			curr[left+right] = '('
			backtrack(left+1, right, curr)
		}

		if right < left {
			curr[left+right] = ')'
			backtrack(left, right+1, curr)
		}
	}

	backtrack(0, 0, make([]byte, n*2))
	return res
}
