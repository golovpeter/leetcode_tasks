package Sliding_window

func maxVowels(s string, k int) int {
	var curRes int

	vowelLetters := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	for i := 0; i < k; i++ {
		if vowelLetters[s[i]] {
			curRes++
		}
	}

	res, l := curRes, 1
	for r := k; r < len(s); r++ {
		if vowelLetters[s[l-1]] {
			curRes--
		}

		if vowelLetters[s[r]] {
			curRes++
		}

		if curRes > res {
			res = curRes
		}

		l++
	}

	return res
}
