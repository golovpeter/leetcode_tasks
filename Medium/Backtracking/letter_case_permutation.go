package Backtracking

import "unicode"

func letterCasePermutation(s string) []string {
	var res []string
	var backtrack func(current, s string, i int)

	backtrack = func(current, s string, i int) {
		if len(current) == len(s) {
			res = append(res, current)
			return
		}

		backtrack(current+string(s[i]), s, i+1)

		runeS := rune(s[i])
		if unicode.IsLetter(runeS) {
			if unicode.IsUpper(runeS) {
				backtrack(current+string(unicode.ToLower(runeS)), s, i+1)
			} else {
				backtrack(current+string(unicode.ToUpper(runeS)), s, i+1)
			}

			return
		}
	}

	backtrack("", s, 0)

	return res
}
