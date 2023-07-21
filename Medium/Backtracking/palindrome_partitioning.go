package Backtracking

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func partition(s string) [][]string {
	var res [][]string
	var backtrack func([]string, string)

	backtrack = func(currCombs []string, s string) {
		if len(s) == 0 {
			res = append(res, append([]string{}, currCombs...))
			return
		}

		for i := 1; i <= len(s); i++ {

			prefix := s[:i]
			postfix := s[i:]

			if isPalindrome(prefix) {
				currCombs = append(currCombs, prefix)
				backtrack(currCombs, postfix)
				currCombs = currCombs[:len(currCombs)-1]
			}
		}

		return
	}

	backtrack([]string{}, s)
	return res
}
