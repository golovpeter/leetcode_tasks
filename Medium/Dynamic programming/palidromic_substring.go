package Dynamic_programming

func countSubstrings(s string) int {
	var res int

	for i := range s {

		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
			res++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
			res++
		}
	}

	return res
}
