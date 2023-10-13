package Backtracking

func maxUniqueSplit(s string) int {
	var maxLen int
	visited := make(map[string]bool)

	var backtrack func(int, int)

	backtrack = func(currResult int, i int) {
		if i == len(s) {
			if currResult > maxLen {
				maxLen = currResult
			}
			return
		}

		for j := i + 1; j <= len(s); j++ {
			if v, ok := visited[s[i:j]]; v && ok {
				continue
			}

			visited[s[i:j]] = true
			backtrack(currResult+1, j)
			visited[s[i:j]] = false
		}
	}

	backtrack(0, 0)
	return maxLen
}
