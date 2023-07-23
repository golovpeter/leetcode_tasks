package Two_pointers

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	l, r := 0, len(needle)-1
	for r < len(haystack) {
		if haystack[l:r+1] == needle {
			return l
		}

		l++
		r++
	}

	return -1

}
