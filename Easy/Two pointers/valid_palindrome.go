package Two_pointers

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !(unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i]))) {
			i++
		}

		for i < j && !(unicode.IsLetter(rune(s[j])) || unicode.IsDigit(rune(s[j]))) {
			j--
		}

		if i == j {
			break
		}

		if rune(s[i]) != rune(s[j]) {
			return false
		}
	}

	return true
}
