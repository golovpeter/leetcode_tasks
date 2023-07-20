package Backtracking

func exist(board [][]byte, word string) bool {
	var backtrack func(board [][]byte, i, j int, word string) bool

	backtrack = func(board [][]byte, i, j int, word string) bool {
		if len(word) == 0 {
			return true
		}

		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[0] {
			return false
		}

		match := false

		originalValue := board[i][j]
		board[i][j] = '_'

		match = backtrack(board, i+1, j, word[1:]) || backtrack(board, i-1, j, word[1:]) ||
			backtrack(board, i, j+1, word[1:]) || backtrack(board, i, j-1, word[1:])

		board[i][j] = originalValue

		return match
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && backtrack(board, i, j, word) {
				return true
			}
		}
	}

	return false
}
