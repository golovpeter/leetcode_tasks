package Sliding_window

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func totalFruit(fruits []int) int {
	if len(fruits) < 3 {
		return len(fruits)
	}

	dict := make(map[int]int, 3)
	start, maxLen := 0, 0

	for end := 0; end < len(fruits); end++ {
		dict[fruits[end]]++

		for len(dict) == 3 {
			dict[fruits[start]]--

			if dict[fruits[start]] == 0 {
				delete(dict, fruits[start])
			}

			start++
		}

		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}
