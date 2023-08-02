package String

func lengthOfLastWord(s string) int {
	counter := 0

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			counter++
		} else if i != len(s)-1 && s[i+1] != ' ' {
			counter = 0
		}
	}

	return counter
}
