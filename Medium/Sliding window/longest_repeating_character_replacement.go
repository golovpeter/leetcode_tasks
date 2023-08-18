package Sliding_window

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}

func characterReplacement(s string, k int) int {
	var maxWindow int
	count := make(map[byte]int)

	l := 0
	for r := 0; r < len(s); r++ {
		count[s[r]]++

		var maxValue int
		for _, value := range count {
			if value > maxValue {
				maxValue = value
			}
		}

		changes := (r - l + 1) - maxValue
		for changes > k {
			count[s[l]]--
			l++
		}

		maxWindow = max(maxWindow, r-l+1)
	}

	return maxWindow
}
