package Backtracking

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var backtrack func(currComb []int, curSum int)

	backtrack = func(currComb []int, curSum int) {
		if curSum > target {
			return
		}

		if curSum == target {
			copyCurrComb := make([]int, len(currComb))
			copy(copyCurrComb, currComb)
			res = append(res, copyCurrComb)
			return
		}

		for _, el := range candidates {
			if el > target || (len(currComb) != 0 && el < currComb[len(currComb)-1]) {
				continue
			}

			currComb = append(currComb, el)
			backtrack(currComb, curSum+el)
			currComb = currComb[:len(currComb)-1]
		}

		return
	}

	backtrack([]int{}, 0)
	return res
}
