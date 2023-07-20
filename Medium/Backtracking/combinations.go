package Backtracking

func combine(n int, k int) [][]int {
	var res [][]int
	var backtrack func(start int, currComb []int)

	backtrack = func(start int, currComb []int) {
		if len(currComb) == k {
			dst := make([]int, len(currComb))
			copy(dst, currComb)
			res = append(res, dst)
			return
		}

		for i := start; i <= n; i++ {
			currComb = append(currComb, i)
			backtrack(i+1, currComb)
			currComb = currComb[:len(currComb)-1]
		}

		return
	}

	backtrack(1, []int{})
	return res
}
