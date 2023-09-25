package Two_pointers

func maxArea(height []int) int {
	var res int
	l, r := 0, len(height)-1

	for l < r {
		ans := (r - l) * min(height[r], height[l])

		if ans > res {
			res = ans
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
