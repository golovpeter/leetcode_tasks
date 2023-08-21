package String

func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		var curSubStr []byte

		for string(curSubStr) != s && len(curSubStr) < len(s) {
			curSubStr = append(curSubStr, []byte(s[:i])...)
		}

		if string(curSubStr) == s {
			return true
		}
	}

	return false
}
