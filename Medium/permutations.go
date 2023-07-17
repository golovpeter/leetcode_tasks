package Medium

func backtrack(currComb, left []int, res *[][]int) {
	if len(left) == 0 {
		*res = append(*res, currComb)
		return
	}

	for idx, _ := range left {
		backtrack(append(currComb, left[idx]),
			append(append([]int{}, left[:idx]...), left[idx+1:]...), res)
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	backtrack(make([]int, 0), nums, &res)
	return res
}
