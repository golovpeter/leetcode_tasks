package Sliding_window

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}

func lengthOfLongestSubstring(s string) int {
	maxSeq := 0
	count := make(map[byte]bool)

	l, r := 0, 0
	for r < len(s) {
		_, ok := count[s[r]]

		if !ok {
			count[s[r]] = true
			r++
		} else {
			maxSeq = max(maxSeq, r-l+1)
			delete(count, s[l])
			l++
		}
	}

	return maxSeq
}
