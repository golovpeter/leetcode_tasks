package Stack

func validateStackSequences(pushed []int, popped []int) bool {
	var j, top int = 0, -1
	for i := 0; i <= len(pushed); i++ {
		for j < len(popped) && top != -1 && pushed[top] == popped[j] {
			top--
			j++
		}
		if i < len(pushed) {
			top++
			pushed[top] = pushed[i]
		}

	}
	return j == len(popped)
}
