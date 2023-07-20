package Backtracking

func combinationSum3(k int, target int) [][]int {
	var res [][]int
	var backtrack func(start int, currComb []int, currSum int)

	backtrack = func(start int, currComb []int, currSum int) {
		if currSum > target || len(currComb) > k {
			return
		}

		if currSum == target && len(currComb) == k {
			res = append(res, append([]int{}, currComb...))
			return
		}

		for i := start; i < 10; i++ {
			currComb = append(currComb, i)
			backtrack(i+1, currComb, currSum+i)
			currComb = currComb[:len(currComb)-1]
		}

		return
	}

	backtrack(1, []int{}, 0)
	return res
}
