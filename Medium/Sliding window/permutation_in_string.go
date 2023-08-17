package Sliding_window

// my solution
//func checkInclusion(s1 string, s2 string) bool {
//	l, r := 0, len(s1)-1
//	freqS2 := make(map[byte]int, len(s1))
//
//	for _, el := range s1 {
//		freqS2[byte(el)]++
//	}
//
//	for r < len(s2) {
//		freq := make(map[byte]int)
//
//		for i := l; i <= r; i++ {
//			freq[s2[i]]++
//		}
//
//		if reflect.DeepEqual(freq, freqS2) {
//			return true
//		}
//
//		l++
//		r++
//	}
//
//	return false
//}

// idiomatic solution
func checkInclusion(s1 string, s2 string) bool {
	l, letters := 0, [26]int{}

	for _, val := range s1 {
		letters[val-97]++
	}

	for r := range s2 {
		letters[s2[r]-97]--

		if letters == [26]int{} {
			return true
		}

		if r+1 >= len(s1) {
			letters[s2[l]-97]++
			l++
		}
	}

	return false
}
