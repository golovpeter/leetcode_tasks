package Arrays

func leastInterval(tasks []byte, n int) int {
	maxFreq := 0
	freq := make(map[byte]int)

	for _, val := range tasks {
		freq[val]++

		if freq[val] > maxFreq {
			maxFreq = freq[val]
		}
	}

	res := (maxFreq - 1) * (n + 1)

	for _, val := range freq {
		if val == maxFreq {
			res++
		}
	}

	return max(res, len(tasks))
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}
